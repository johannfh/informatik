import * as ex from "excalibur";

// They contain a bunch of useful components that you might use
// actor.transform
// actor.motion
// actor.graphics
// actor.body
// actor.collider
// actor.actions
// actor.pointer

export class Button extends ex.Actor {
    constructor(
        position: ex.Vector | [number, number],
        public action: () => void,
        public actionName: string,
        public sprite: ex.Sprite,
    ) {
        let pos = position instanceof ex.Vector ? position : ex.vec(position[0], position[1]);

        super({
            name: `Button-${actionName}`,
            pos: pos,
            width: 100,
            height: 100,
        });

        this.sprite.scale = this.sprite.scale.normalize();
        this.sprite.scale.x /= 5;
        this.sprite.scale.y /= 5;
    }

    override onInitialize() {
        this.graphics.add(this.sprite);
    }
}
