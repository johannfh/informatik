import * as ex from "excalibur";
import { MAX_X, MAX_Y, MIN_X, MIN_Y } from "./constants";

export type AnimalArgs = {
    position: ex.Vector;
    speed: number;
} & ex.ActorArgs;

export abstract class Animal extends ex.Actor {
    private _targetPosition: ex.Vector;
    public speed: number;

    public get targetPosition(): ex.Vector {
        return this._targetPosition;
    }

    public set targetPosition(pos: ex.Vector) {
        this._targetPosition.x = ex.clamp(pos.x, MIN_X, MAX_X);
        this._targetPosition.y = ex.clamp(pos.y, MIN_Y, MAX_Y);
    }

    constructor(args: AnimalArgs) {
        args.width ??= 20;
        args.height ??= 20;
        args.pos ??= args.position;
        super(args);

        this._targetPosition = args.position;
        this.speed = args.speed;
    }

    override onPostUpdate(_: ex.Engine, elapsedMs: number): void {
        let deltatime = elapsedMs / 1000;

        let direction = this.targetPosition.sub(this.pos).normalize();
        this.pos = this.pos.add(direction.scale(this.speed * deltatime));

        this.movementLogic(deltatime);
        this.hunger(elapsedMs);
    }

    public abstract movementLogic(deltatime: number): void;
    public abstract hunger(elapsedMs: number): void;

    public wander(wanderDistance: number) {
        this.targetPosition = this.pos.add(
            ex
                .vec(Math.random() - 0.5, Math.random() - 0.5)
                .normalize()
                .scale(wanderDistance),
        );
    }

    public runFrom(position: ex.Vector, runDistance: number) {
        this.targetPosition = this.pos.add(this.pos.sub(position).normalize().scale(runDistance));
    }
}
