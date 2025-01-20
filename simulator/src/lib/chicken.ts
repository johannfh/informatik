import * as ex from "excalibur";
import { Animal, type AnimalArgs } from "./animal";
import { Resources } from "./resources";
import type { ActorGetter } from "./types";
import { Fox } from "./fox";
import { closestActorTo } from "./utils";

export type ChickenArgs = {
    actorGetter: ActorGetter;
} & AnimalArgs;

export class Chicken extends Animal {
    actorGetter: ActorGetter;
    sprite: ex.Sprite;

    constructor(args: ChickenArgs) {
        super(args);
        this.actorGetter = args.actorGetter;
        let sprite = (this.sprite = Resources.Chicken.toSprite());
        sprite.scale = sprite.scale.normalize();
        sprite.scale.x /= 5;
        sprite.scale.y /= 5;
    }

    override onInitialize(engine: ex.Engine) {
        super.onInitialize(engine);

        this.graphics.add(this.sprite);
    }

    private getPredators(animalGetter: ActorGetter): Animal[] {
        return animalGetter((a) => a instanceof Fox) as Animal[];
    }

    private getClosestPredator(animalGetter: ActorGetter): Animal | undefined {
        const predators = this.getPredators(animalGetter);
        if (predators.length === 0) return undefined;
        return closestActorTo(this.pos, predators);
    }

    override hunger(elapsedMs: number): void {
        // Chicken = Food
    }

    public movementLogic(deltatime: number): void {
        if (this.pos.distance(this.targetPosition) > 1) {
            return;
        }

        let predator = this.getClosestPredator(this.actorGetter);

        if (!predator) {
            this.wander(50);
            return;
        }

        this.runFrom(predator.pos, 100);
    }
}
