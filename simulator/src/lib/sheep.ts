import * as ex from "excalibur";
import { Animal, type AnimalArgs } from "./animal";
import { Resources } from "./resources";
import type { ActorGetter } from "./types";
import { Wolf } from "./wolf";
import { closestActorTo } from "./utils";

export type SheepArgs = {
    actorGetter: ActorGetter;
} & AnimalArgs;

export class Sheep extends Animal {
    private actorGetter: ActorGetter;
    sprite: ex.Sprite;

    constructor(args: SheepArgs) {
        super(args);
        this.actorGetter = args.actorGetter;
        let sprite = (this.sprite = Resources.Sheep.toSprite());
        sprite.scale = sprite.scale.normalize();
        sprite.scale.x /= 5;
        sprite.scale.y /= 5;
    }

    override onInitialize(engine: ex.Engine) {
        super.onInitialize(engine);

        this.graphics.add(this.sprite);
    }

    private getPredators(actorGetter: ActorGetter): Animal[] {
        return actorGetter((a) => a instanceof Wolf) as Animal[];
    }

    private getClosestPredator(actorGetter: ActorGetter): Animal | undefined {
        const predators = this.getPredators(actorGetter);
        if (predators.length === 0) return undefined;
        return closestActorTo(this.pos, predators);
    }

    override hunger(elapsedMs: number): void {
        // Sheep = Food
    }

    public movementLogic(deltatime: number): void {
        if (this.pos.distance(this.targetPosition) > 1) {
            return;
        }

        let predator = this.getClosestPredator(this.actorGetter);

        if (!predator || this.pos.distance(predator.pos) > 100) {
            this.wander(50);
            return;
        }

        this.runFrom(predator.pos, 100);
    }
}
