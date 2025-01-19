import * as ex from "excalibur";
import { PredatorGetter, Prey } from "./prey";
import { Predator, PreyGetter } from "./predator";
import { Resources } from "./resources";

function CreatePreyCreator(
    speed: number,
    name?: string,
    predatorGetter?: PredatorGetter,
    sprite?: ex.Sprite,
) {
    return () =>
        new Prey(
            speed,
            name,
            ex.vec((Math.random() - 0.5) * 200 + 500, 600),
            predatorGetter,
            sprite,
        );
}

function CreatePredatorCreator(
    speed: number,
    name?: string,
    preyGetter?: PreyGetter,
    sprite?: ex.Sprite,
) {
    return () =>
        new Predator(
            speed,
            name,
            ex.vec((Math.random() - 0.5) * 200 + 500, 400),
            preyGetter,
            sprite,
        );
}

export class MainScene extends ex.Scene {
    private chickenName = "chicken";
    private chickenCreator: () => Prey;
    private chickenPredatorGetter: PredatorGetter;

    private foxName = "fox";
    private foxCreator: () => Predator;
    private foxPreyGetter: PreyGetter;

    private sheepName = "sheep";
    private sheepCreator: () => Prey;
    private sheepPredatorGetter: PredatorGetter;

    private wolfName = "wolf";
    private wolfCreator: () => Predator;
    private wolfPreyGetter: PreyGetter;

    constructor() {
        super();
        let thisScene = this;

        this.chickenPredatorGetter = () =>
            thisScene.actors.filter(
                (actor) => actor.name === thisScene.foxName,
            );
        this.chickenCreator = CreatePreyCreator(
            50,
            thisScene.chickenName,
            this.chickenPredatorGetter,
            Resources.Chicken.toSprite(),
        );

        this.foxPreyGetter = () =>
            thisScene.actors.filter(
                (actor) => actor.name === thisScene.chickenName,
            );
        this.foxCreator = CreatePredatorCreator(
            150,
            thisScene.foxName,
            this.foxPreyGetter,
            Resources.Fox.toSprite(),
        );

        this.sheepPredatorGetter = () =>
            thisScene.actors.filter(
                (actor) => actor.name === thisScene.wolfName,
            );
        this.sheepCreator = CreatePreyCreator(
            30,
            thisScene.sheepName,
            this.sheepPredatorGetter,
            Resources.Sheep.toSprite(),
        );

        this.wolfPreyGetter = () =>
            thisScene.actors.filter(
                (actor) => actor.name === thisScene.sheepName,
            );
        this.wolfCreator = CreatePredatorCreator(
            70,
            thisScene.wolfName,
            this.wolfPreyGetter,
            Resources.Wolf.toSprite(),
        );
    }

    override onInitialize(engine: ex.Engine): void {
        const chicken = this.chickenCreator();
        this.add(chicken);
        const fox = this.foxCreator();
        this.add(fox);

        const sheep = this.sheepCreator();
        this.add(sheep);
        const wolf = this.wolfCreator();
        this.add(wolf);
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
