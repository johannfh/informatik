export type AnimalGetter = (names: AnimalOption[]) => ex.Actor[];

export type AnimalOption = PreyAnimal | PredatorAnimal;

export const PredatorAnimals = ["wolf", "fox"] as const;
export type PredatorAnimal = (typeof PredatorAnimals)[number];

export const PreyAnimals = ["sheep", "chicken"] as const;
export type PreyAnimal = (typeof PreyAnimals)[number];
