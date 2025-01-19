import * as ex from "excalibur";
import { Resources } from "./resources";
import { screenHeight, screenWidth } from "./constants";

// They contain a bunch of useful components that you might use
// actor.transform
// actor.motion
// actor.graphics
// actor.body
// actor.collider
// actor.actions
// actor.pointer

export type PredatorGetter = () => ex.Actor[];

const MIN_X = 10;
const MAX_X = screenWidth - 10;

const MIN_Y = 10;
const MAX_Y = screenHeight - 10;

export class Prey extends ex.Actor {
  private _targetPos: ex.Vector;

  private get targetPos(): ex.Vector {
    return this._targetPos;
  }

  private set targetPos(pos: ex.Vector) {
    this._targetPos.x = ex.clamp(pos.x, MIN_X, MAX_X);
    this._targetPos.y = ex.clamp(pos.y, MIN_Y, MAX_Y);
  }

  constructor(
    public speed: number,
    name?: string,
    pos: ex.Vector = ex.vec(500, 500),
    private predatorGetter: PredatorGetter = () => [],
    private sprite: ex.Sprite = Resources.Sword.toSprite(),
  ) {
    super({
      name: name,
      pos: pos,
      width: 100,
      height: 100,
    });

    this._targetPos = pos;

    this.sprite.scale = this.sprite.scale.normalize();
    this.sprite.scale.x /= 5;
    this.sprite.scale.y /= 5;
  }

  override onInitialize() {
    this.graphics.add(this.sprite);
  }

  override onPostUpdate(engine: ex.Engine, elapsedMs: number): void {
    let deltatime = elapsedMs / 1000;

    // Put any update logic here runs every frame after Actor builtins
    this.movementLogic(deltatime);
  }

  private movementLogic(deltatime: number) {
    let enemies = this.predatorGetter();
    let speed = 100;

    if (this.pos.distance(this.targetPos) > 1) {
      let direction = this.targetPos.sub(this.pos).normalize();
      this.pos = this.pos.add(direction.scale(speed * deltatime));
      return;
    }

    // wander when no predator exists
    if (enemies.length === 0) {
      this.wander();
      return;
    }

    // initialize the closestEnemy object
    let closestEnemy = {
      distance: this.pos.distance(enemies[0].pos),
      actor: enemies[0],
    };

    // calculate the actual closestEnemy
    for (const actor of enemies) {
      let distance = actor.pos.distance(this.pos);
      if (distance < closestEnemy.distance) {
        closestEnemy = { distance, actor };
      }
    }

    // TODO: just wander when enemy out of `this.sightRadius`

    this.runFrom(closestEnemy.actor.pos);
  }

  private wander() {
    let wanderDistance = 100;

    this.targetPos = this.pos.add(
      ex.vec(
        (Math.random() - 0.5) * 2 * wanderDistance,
        (Math.random() - 0.5) * 2 * wanderDistance,
      ),
    );
    console.log(
      `from ${this.pos.x.toFixed(3)}-${this.pos.y.toFixed(3)} to ${this.targetPos.x.toFixed(3)}-${this.targetPos.y.toFixed(3)}`,
    );
  }

  private runFrom(position: ex.Vector) {
    let runDistance = 100;
    this.targetPos = this.pos.add(
      this.pos.sub(position).normalize().scale(runDistance),
    );
  }
}
