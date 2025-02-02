import * as ex from "excalibur";
import { Resources } from "./resources";

// Actors are the main unit of composition you'll likely use, anything that you want to draw and move around the screen
// is likely built with an actor

// They contain a bunch of useful components that you might use
// actor.transform
// actor.motion
// actor.graphics
// actor.body
// actor.collider
// actor.actions
// actor.pointer

export class Player extends ex.Actor {
    constructor() {
        super({
            // Giving your actor a name is optional, but helps in debugging using the dev tools or debug mode
            // https://github.com/excaliburjs/excalibur-extension/
            // Chrome: https://chromewebstore.google.com/detail/excalibur-dev-tools/dinddaeielhddflijbbcmpefamfffekc
            // Firefox: https://addons.mozilla.org/en-US/firefox/addon/excalibur-dev-tools/
            name: "Player",
            pos: ex.vec(150, 150),
            width: 100,
            height: 100,
            // anchor: vec(0, 0), // Actors default center colliders and graphics with anchor (0.5, 0.5)
            // collisionType: CollisionType.Active, // Collision Type Active means this participates in collisions read more https://excaliburjs.com/docs/collisiontypes
        });
    }

    override onInitialize() {
        // Generally recommended to stick logic in the "On initialize"
        // This runs before the first update
        // Useful when
        // 1. You need things to be loaded like Images for graphics
        // 2. You need excalibur to be initialized & started
        // 3. Deferring logic to run time instead of constructor time
        // 4. Lazy instantiation
        this.graphics.add(Resources.Sword.toSprite());

        // Actions are useful for scripting common behavior, for example patrolling enemies
        this.actions.repeatForever((ctx) => {});

        // Sometimes you want to click on an actor!
        this.on("pointerdown", (evt) => {
            // Pointer events tunnel in z order from the screen down, you can cancel them!
            // if (true) {
            //   evt.cancel();
            // }
            console.log("You clicked the actor @", evt.worldPos.toString());
        });
    }

    override onPreUpdate(engine: ex.Engine, elapsedMs: number): void {
        // Put any update logic here runs every frame before Actor builtins
    }

    override onPostUpdate(engine: ex.Engine, elapsedMs: number): void {
        // Put any update logic here runs every frame after Actor builtins
    }

    override onPreCollisionResolve(
        self: ex.Collider,
        other: ex.Collider,
        side: ex.Side,
        contact: ex.CollisionContact,
    ): void {
        // Called before a collision is resolved, if you want to opt out of this specific collision call contact.cancel()
    }

    override onPostCollisionResolve(
        self: ex.Collider,
        other: ex.Collider,
        side: ex.Side,
        contact: ex.CollisionContact,
    ): void {
        // Called every time a collision is resolved and overlap is solved
    }

    override onCollisionStart(
        self: ex.Collider,
        other: ex.Collider,
        side: ex.Side,
        contact: ex.CollisionContact,
    ): void {
        // Called when a pair of objects are in contact
    }

    override onCollisionEnd(
        self: ex.Collider,
        other: ex.Collider,
        side: ex.Side,
        lastContact: ex.CollisionContact,
    ): void {
        // Called when a pair of objects separates
    }
}
