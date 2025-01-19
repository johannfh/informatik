import * as ex from "excalibur";

// It is convenient to put your resources in one place
export const Resources = {
  Sword: new ex.ImageSource("./images/sword.png"),

  Chicken: new ex.ImageSource("./images/chicken.png"),
  Fox: new ex.ImageSource("./images/fox.png"),

  Sheep: new ex.ImageSource("./images/sheep.png"),
  Wolf: new ex.ImageSource("./images/wolf.png"),
} as const;

// We build a loader and add all of our resources to the boot loader
// You can build your own loader by extending DefaultLoader
export const loader = new ex.Loader();
for (const res of Object.values(Resources)) {
  loader.addResource(res);
}
