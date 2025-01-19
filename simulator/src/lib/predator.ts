import * as ex from "excalibur";
import { Resources } from "./resources";
import { screenHeight, screenWidth } from "./constants";
import type { AnimalGetter, AnimalOption, PredatorAnimal } from "./types";
import type { PreyAnimal } from "./types";

export type PreyGetter = () => ex.Actor[];

const MIN_X = 10;
const MAX_X = screenWidth - 10;

const MIN_Y = 10;
const MAX_Y = screenHeight - 10;

export type NewPredatorParams = {
  type: PredatorAnimal;
  position: ex.Vector;
  animalGetter: AnimalGetter;
};

export class Predator extends ex.Actor {
  private _targetPos: ex.Vector;
  private speed: number;
  private animalGetter: AnimalGetter;
  sprite: ex.Sprite;
  preyList: AnimalOption[];

  private get targetPos(): ex.Vector {
    return this._targetPos;
  }

  private set targetPos(pos: ex.Vector) {
    this._targetPos.x = ex.clamp(pos.x, MIN_X, MAX_X);
    this._targetPos.y = ex.clamp(pos.y, MIN_Y, MAX_Y);
  }

  constructor({ position, animalGetter, type }: NewPredatorParams) {
    super({
      name: type,
      pos: position,
      width: 20,
      height: 20,
    });

    this._targetPos = position;

    if (type === "fox") {
      // fox
      this.sprite = Resources.Fox.toSprite();
      this.speed = 120;
      this.preyList = ["chicken"];
    } else {
      // wolf
      this.sprite = Resources.Wolf.toSprite();
      this.speed = 70;
      this.preyList = ["sheep"];
    }

    this.animalGetter = animalGetter;

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

  override onCollisionStart(
    self: ex.Collider,
    other: ex.Collider,
    side: ex.Side,
    contact: ex.CollisionContact,
  ): void {
    if (!this.preyList.includes(other.owner.name as PreyAnimal)) return;
    other.owner.kill();
  }

  private movementLogic(deltatime: number) {
    let enemies = this.animalGetter(this.preyList);
    if (this.pos.distance(this.targetPos) > 1) {
      let direction = this.targetPos.sub(this.pos).normalize();
      this.pos = this.pos.add(direction.scale(this.speed * deltatime));
      return;
    }

    // wander when no prey exists
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

    let attackDistance = 30;
    this.targetPos = this.pos.add(
      closestEnemy.actor.pos
        .sub(this.pos)
        .normalize()
        .rotate(((Math.random() - 0.5) * Math.PI) / 1.5)
        .scale(attackDistance),
    );
  }

  private wander() {
    let wanderDistance = 100;

    this.targetPos = this.pos.add(
      ex.vec(
        (Math.random() - 0.5) * 2 * wanderDistance,
        (Math.random() - 0.5) * 2 * wanderDistance,
      ),
    );
  }
}
