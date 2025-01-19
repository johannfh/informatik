import * as ex from "excalibur";
import { loader } from "./resources";
import { MainScene } from "./main-scene";
import { screenHeight, screenWidth } from "./constants";

// Goal is to keep main.ts small and just enough to configure the engine

const game = new ex.Engine({
    width: screenWidth, // Logical width and height in game pixels
    height: screenHeight,
    displayMode: ex.DisplayMode.FitScreen, // Display mode tells excalibur how to fill the window
    pixelArt: true, // pixelArt will turn on the correct settings to render pixel art without jaggies or shimmering artifacts
    scenes: {
        start: MainScene,
    },
    // physics: {
    //   solver: SolverStrategy.Realistic,
    //   substep: 5 // Sub step the physics simulation for more robust simulations },
    // fixedUpdateTimestep: 16 // Turn on fixed update timestep when consistent physic simulation is important
});

game.start("start", {
    // name of the start scene 'start'
    loader, // Optional loader (but needed for loading images/sounds)
    inTransition: new ex.FadeInOut({
        // Optional in transition
        duration: 200,
        direction: "in",
        color: ex.Color.ExcaliburBlue,
    }),
}).then(() => {
    // Do something after the game starts
    console.log("game started");
});
