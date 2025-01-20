import * as ex from "excalibur";
import { Resources } from "./resources";
import type { ActorGetter } from "./types";
import { Animal, type AnimalArgs } from "./animal";
import { Chicken } from "./chicken";
import { FOX_COOLDOWN, FOX_SATURATION } from "./constants";
import { closestActorTo } from "./utils";

export type FoxArgs = {
    actorGetter: ActorGetter;
} & AnimalArgs;

export class Fox extends Animal {
    private sprite: ex.Sprite;
    private actorGetter: ActorGetter;
    private saturation = FOX_SATURATION;
    cooldown: number = 0;

    constructor(params: FoxArgs) {
        super(params);
        this.actorGetter = params.actorGetter;

        let sprite = (this.sprite = Resources.Fox.toSprite());
        sprite.scale = sprite.scale.normalize();
        sprite.scale.x /= 5;
        sprite.scale.y /= 5;
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
        if (other.owner instanceof Chicken && this.cooldown < 0) {
            other.owner.kill();
            this.saturation = FOX_SATURATION;
        }
    }

    private getPrey(actorGetter: ActorGetter): Animal[] {
        return actorGetter((a) => a instanceof Chicken) as Animal[];
    }

    private getClosestPrey(actorGetter: ActorGetter): Animal | undefined {
        const prey = this.getPrey(actorGetter);
        if (prey.length === 0) return undefined;
        return closestActorTo(this.pos, prey)
    }

    override hunger(elapsedMs: number) {
        if (this.cooldown < 0) this.saturation -= elapsedMs;
        if (this.saturation < 0) {
            this.kill();
        }
    }

    override movementLogic(deltatime: number) {
        if (this.pos.distance(this.targetPosition) > 1) {
            return;
        }

        const prey = this.getClosestPrey(this.actorGetter);

        // wander when no prey exists
        if (!prey) {
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
                .rotate(((Math.random() - 0.5) * Math.PI) / 1.5)
                // scale walk distance by attack distance
                .scale(attackDistance),
        );
    }
}
