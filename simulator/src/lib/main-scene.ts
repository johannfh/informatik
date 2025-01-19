import * as ex from "excalibur";
import { writable, type Writable } from "svelte/store";
import type { ActorGetter, AnimalOption } from "./types";
import { Sheep } from "./sheep";
import { Wolf } from "./wolf";
import { Fox } from "./fox";
import { Chicken } from "./chicken";

export type SpawnAnimalEvent = {
    type: "spawnAnimalEvent";
    animal: AnimalOption;
    position: ex.Vector;
};
export type KillAnimalEvent = {
    type: "killAnimalEvent";
    position: ex.Vector;
};

export type InputEvent = SpawnAnimalEvent | KillAnimalEvent;

export let inputEvents: Writable<InputEvent[]> = writable([]);

export class MainScene extends ex.Scene {
    private actorGetter: ActorGetter = (func) => this.actors.filter((a) => func(a));

    override onInitialize(engine: ex.Engine): void {
        inputEvents.subscribe((v) => {
            let event = v.shift();
            if (!event) return;
            console.log(`handling ${event.type}`);
            this.handleInputEvent(event);
        });
    }

    private handleInputEvent(event: InputEvent) {
        if (event.type === "spawnAnimalEvent") {
            this.handleSpawnAnimal(event);
            return;
        }

        console.error(`event handler for '${event.type}' currently unimplemented`);
    }

    private handleSpawnAnimal({ animal, position }: SpawnAnimalEvent) {
        const actorGetter = this.actorGetter;

        switch (animal) {
            case Wolf:
                this.add(new Wolf({ position, speed: 50, actorGetter }));
                break;

            case Sheep:
                this.add(new Sheep({ position, speed: 50, actorGetter }));
                break;

            case Fox:
                this.add(new Fox({ position, speed: 50, actorGetter }));
                break;

            case Chicken:
                this.add(new Chicken({ position, speed: 50, actorGetter }));
                break;
        }
    }

    override onPreLoad(loader: ex.DefaultLoader): void {
        // Add any scene specific resources to load
    }

    override onActivate(context: ex.SceneActivationContext<unknown>): void {
        // Called when Excalibur transitions to this scene
        // Only 1 scene is active at a time
    }

    override onDeactivate(context: ex.SceneActivationContext): void {
        // Called when Excalibur transitions away from this scene
        // Only 1 scene is active at a time
    }

    override onPreUpdate(engine: ex.Engine, elapsedMs: number): void {
        // Called before anything updates in the scene
    }

    override onPostUpdate(engine: ex.Engine, elapsedMs: number): void {
        // Called after everything updates in the scene
    }

    override onPreDraw(ctx: ex.ExcaliburGraphicsContext, elapsedMs: number): void {
        // Called before Excalibur draws to the screen
    }

    override onPostDraw(ctx: ex.ExcaliburGraphicsContext, elapsedMs: number): void {
        // Called after Excalibur draws to the screen
    }
}
