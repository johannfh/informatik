<script lang="ts">
    import { onMount } from "svelte";
    import * as ex from "excalibur";
    import { loader } from "$lib/resources";
    import { MainScene } from "$lib/main-scene";
    import { screenHeight, screenWidth } from "$lib/constants";

    let gameCanvas: HTMLCanvasElement;

    onMount(() => {
        let game = new ex.Engine({
            width: screenWidth,
            height: screenHeight,
            canvasElement: gameCanvas,
            scenes: {
                start: MainScene,
            },
        })

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
    })
</script>

<div class="flex flex-col">
    <div class="menu">

    </div>
    <div class="p-4">
        <canvas bind:this={gameCanvas}></canvas>
    </div>
</div>