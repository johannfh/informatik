import * as ex from "excalibur";
import { Animal, type AnimalArgs } from "./animal";
import { Resources } from "./resources";
import type { ActorGetter } from "./types";
import { Sheep } from "./sheep";
import { closestActorTo } from "./utils";
import { WOLF_SATURATION } from "./constants";

export type WolfArgs = {
    actorGetter: ActorGetter;
} & AnimalArgs;

export class Wolf extends Animal {
    private actorGetter: ActorGetter;
    sprite: ex.Sprite;
    cooldown: number;
    saturation: number;

    constructor(args: WolfArgs) {
        super(args);
        this.actorGetter = args.actorGetter;
        let sprite = (this.sprite = Resources.Wolf.toSprite());
        sprite.scale = sprite.scale.normalize();
        sprite.scale.x /= 5;
        sprite.scale.y /= 5;
        this.cooldown = 0;
        this.saturation = WOLF_SATURATION;
    }

    override onInitialize(engine: ex.Engine) {
        super.onInitialize(engine);

        this.graphics.add(this.sprite);
    }

    override onPostUpdate(e: ex.Engine, elapsedMs: number): void {
        super.onPostUpdate(e, elapsedMs);
        this.cooldown -= elapsedMs;
    }

    override onCollisionStart(
        self: ex.Collider,
        other: ex.Collider,
        side: ex.Side,
        contact: ex.CollisionContact,
    ): void {
        super.onCollisionStart(self, other, side, contact);
        if (other.owner instanceof Sheep && this.cooldown < 0) {
            other.owner.kill();
            this.cooldown = 3000;
        }
    }

    override hunger(elapsedMs: number): void {
        if (this.cooldown < 0) this.saturation -= elapsedMs;
        if (this.saturation < 0) {
            this.kill()
        }
    }

    private getPrey(actorGetter: ActorGetter): Animal[] {
        return actorGetter((a) => a instanceof Sheep) as Animal[];
    }

    private getClosestPrey(actorGetter: ActorGetter): Animal | undefined {
        const prey = this.getPrey(actorGetter);
        if (prey.length === 0) return undefined;
        return closestActorTo(this.pos, prey)
    }

    public movementLogic(deltatime: number): void {
        if (this.pos.distance(this.targetPosition) > 1) {
            return;
        }

        const prey = this.getClosestPrey(this.actorGetter);

        if (!prey || this.cooldown > 0) {
            this.wander(50);
            return;
        }

        const attackDistance = 30;

        this.targetPosition = this.pos.add(
            prey.pos
                // calculate direction to prey
                .sub(this.pos)
                .normalize()
                // randomize direction to prey
                .rotate(((Math.random() - 0.5) * Math.PI) / 3)
                // scale walk distance by attack distance
                .scale(attackDistance),
        );
    }
}
