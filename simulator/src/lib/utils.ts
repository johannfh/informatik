import * as ex from "excalibur";

export function closestActorTo<T extends ex.Actor>(position: ex.Vector, actors: T[]) {
    // initialize the closestActor object
    let closestActor = {
        distance: position.distance(actors[0].pos),
        actor: actors[0],
    };

    // calculate the actual closestEnemy
    for (const actor of actors) {
        let distance = actor.pos.distance(position);

        if (distance < closestActor.distance) {
            closestActor = { distance, actor };
        }
    }

    return closestActor.actor;
}
