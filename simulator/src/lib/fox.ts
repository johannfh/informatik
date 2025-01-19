import * as ex from "excalibur";
import { Resources } from "./resources";
import type { ActorGetter } from "./types";
import { Animal, type AnimalArgs } from "./animal";
import { Chicken } from "./chicken";

export type FoxArgs = {
    actorGetter: ActorGetter;
} & AnimalArgs;

export class Fox extends Animal {
    private sprite: ex.Sprite;
    private actorGetter: ActorGetter;

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

    override onCollisionStart(
        self: ex.Collider,
        other: ex.Collider,
        side: ex.Side,
        contact: ex.CollisionContact,
    ): void {
        super.onCollisionStart(self, other, side, contact);
        if (other.owner instanceof Chicken) {
            other.owner.kill();
        }
    }

    private getPrey(actorGetter: ActorGetter): Animal[] {
        return actorGetter((a) => a instanceof Chicken) as Animal[];
    }

    private getClosestPrey(actorGetter: ActorGetter): Animal | undefined {
        const prey = this.getPrey(actorGetter);

        if (prey.length === 0) return undefined;

        // initialize the closestEnemy object
        let closestPrey = {
            distance: this.pos.distance(prey[0].pos),
            actor: prey[0],
        };

        // calculate the actual closestEnemy
        for (const actor of prey) {
            let distance = actor.pos.distance(this.pos);

            if (distance < closestPrey.distance) {
                closestPrey = { distance, actor };
            }
        }

        return closestPrey.actor;
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
