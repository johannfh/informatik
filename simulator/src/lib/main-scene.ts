import * as ex from "excalibur";
import { type PredatorGetter, Prey } from "./prey";
import { Predator, type PreyGetter } from "./predator";
import { Resources } from "./resources";
import { screenWidth } from "./constants";
import { writable, type Writable } from "svelte/store";

const CreatePreyCreator = (
  speed: number,
  name?: string,
  predatorGetter?: PredatorGetter,
  sprite?: ex.Sprite,
) => (position: ex.Vector) =>
    new Prey(
      speed,
      name,
      position,
      predatorGetter,
      sprite,
    );


const CreatePredatorCreator = (
  speed: number,
  name?: string,
  preyGetter?: PreyGetter,
  sprite?: ex.Sprite,
) => (position: ex.Vector) => new Predator(speed, name, position, preyGetter, sprite);


export type AnimalOption = "wolf" | "sheep" | "fox" | "chicken";
export type SpawnAnimalEvent = {
  type: "spawnAnimalEvent",
  animal: AnimalOption;
  position: ex.Vector;
};
export type KillAnimalEvent = {
  type: "killAnimalEvent",
  position: ex.Vector;
};

export type InputEvent = SpawnAnimalEvent | KillAnimalEvent;

export let inputEvents: Writable<InputEvent[]> = writable([]);

export class MainScene extends ex.Scene {
  private chickenName = "chicken";
  private chickenCreator: (position: ex.Vector) => Prey;
  private chickenPredatorGetter: PredatorGetter;

  private foxName = "fox";
  private foxCreator: (position: ex.Vector) => Predator;
  private foxPreyGetter: PreyGetter;

  private sheepName = "sheep";
  private sheepCreator: (position: ex.Vector) => Prey;
  private sheepPredatorGetter: PredatorGetter;

  private wolfName = "wolf";
  private wolfCreator: (position: ex.Vector) => Predator;
  private wolfPreyGetter: PreyGetter;

  constructor() {
    super();
    let thisScene = this;

    this.chickenPredatorGetter = () =>
      thisScene.actors.filter((actor) => actor.name === thisScene.foxName);
    this.chickenCreator = CreatePreyCreator(
      50,
      thisScene.chickenName,
      this.chickenPredatorGetter,
      Resources.Chicken.toSprite(),
    );

    this.foxPreyGetter = () =>
      thisScene.actors.filter((actor) => actor.name === thisScene.chickenName);
    this.foxCreator = CreatePredatorCreator(
      150,
      thisScene.foxName,
      this.foxPreyGetter,
      Resources.Fox.toSprite(),
    );

    this.sheepPredatorGetter = () =>
      thisScene.actors.filter((actor) => actor.name === thisScene.wolfName);
    this.sheepCreator = CreatePreyCreator(
      30,
      thisScene.sheepName,
      this.sheepPredatorGetter,
      Resources.Sheep.toSprite(),
    );

    this.wolfPreyGetter = () =>
      thisScene.actors.filter((actor) => actor.name === thisScene.sheepName);
    this.wolfCreator = CreatePredatorCreator(
      70,
      thisScene.wolfName,
      this.wolfPreyGetter,
      Resources.Wolf.toSprite(),
    );
  }

  override onInitialize(engine: ex.Engine): void {
    inputEvents.subscribe(v => {
      let event = v.shift()
      if (!event) return;

      switch (event.type) {
        case "spawnAnimalEvent":
          this.handleSpawnAnimal(event)
          break;

        default:
          console.error(`event '${event.type}' currently unimplemented`)
          break;
      
      }
    })
  }
  
  private handleSpawnAnimal({ animal, position }: SpawnAnimalEvent) {
    switch (animal) {
      case "fox":
        this.add(this.foxCreator(position))
        break;
      case "chicken":
        this.add(this.chickenCreator(position))
        break;
      case "wolf":
        this.add(this.wolfCreator(position))
        break;
      case "sheep":
        this.add(this.sheepCreator(position))
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
