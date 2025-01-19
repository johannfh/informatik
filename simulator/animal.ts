import { Actor, Vector } from "excalibur";

export class Animal extends Actor {
    constructor(pos?: Vector) {
        super({ pos });
    }
}
