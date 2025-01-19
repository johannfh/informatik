import * as ex from "excalibur";
import { Prey } from "./prey";
import { type PreyAnimal } from "./types";
import { PreyAnimals } from "./types";
import { Predator } from "./predator";
import { writable, type Writable } from "svelte/store";
import type { AnimalOption } from "./types";

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
  private animalGetter = (names: AnimalOption[]): ex.Actor[] => {
    return this.actors.filter((v) => (names as string[]).includes(v.name));
  };

  override onInitialize(engine: ex.Engine): void {
    inputEvents.subscribe((v) => {
      let event = v.shift();
      if (!event) return;
      console.log(`handling ${event.type}`);

      switch (event.type) {
        case "spawnAnimalEvent":
          this.handleSpawnAnimal(event);
          break;

        default:
          console.error(
            `event handler for '${event.type}' currently unimplemented`,
          );
          break;
      }
    });
  }

  private handleSpawnAnimal({ animal, position }: SpawnAnimalEvent) {
    const type = animal;
    const animalGetter = this.animalGetter;

    if (type === "fox" || type === "wolf") {
      this.add(new Predator({ position, type, animalGetter }));
    } else {
      this.add(new Prey({ position, type, animalGetter }));
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

  override onPreDraw(
    ctx: ex.ExcaliburGraphicsContext,
    elapsedMs: number,
  ): void {
    // Called before Excalibur draws to the screen
  }

  override onPostDraw(
    ctx: ex.ExcaliburGraphicsContext,
    elapsedMs: number,
  ): void {
    // Called after Excalibur draws to the screen
  }
}
