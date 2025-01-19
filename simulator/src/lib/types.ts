import * as ex from "excalibur";
import { Fox } from "./fox";
import { Wolf } from "./wolf";
import { Sheep } from "./sheep";
import { Chicken } from "./chicken";

export type ActorFilterFunc = (actor: ex.Actor) => boolean;
export type ActorGetter = (filterFunc: ActorFilterFunc) => ex.Actor[];

export type AnimalOption = PreyAnimal | PredatorAnimal;

export const PredatorAnimals = [Wolf, Fox] as const;
export type PredatorAnimal = (typeof PredatorAnimals)[number];

export const PreyAnimals = [Sheep, Chicken] as const;
export type PreyAnimal = (typeof PreyAnimals)[number];
