export default {
  title: "Flippour",
  link: "flippour",
  body: `
    A simple puzzle-esque mobile game for Android and iOS
    <br/>
    <br/>
    The game is very simple. The user is presented with a color and has to tap all squares
    of that color before the time runs out. Each square correctly tapped is +1 while a square
    incorrectly tapped is -10. If the score reaches below zero the game is ended. If the timer 
    runs out the game is ended. Each ten levels the totaly level time it dropped by one seconds,
    i.e. levels 1 - 10 the user has ten seconds to complete the level, while levels 11 - 20 the user
    has nine seconds to complete the level.
    <br/>
    <br/>
    This project implements React Native's Animated API along with the option "useNativeDriver" for using
    animation on the native level to ensure the user does not experience any jank on the main UI thread along
    with the JS thread. High performance was a big piece of this project and it's working beautifully.
    <br/>
    <br/>
    Originially this project was mean to try out Facebook and Expo's "create-react-native-app" platform, but
    has since been abonded in favor of Facebook's original react-native CLI.
    <br/>
    <br/>
    You can find Play Store and App Store links at <a href="https://mini-eggs.github.io/Flippour/">here</a>.
    <br/>
    <br/>
    This project has been open sourced. You can find the code <a href="https://github.com/mini-eggs/Flippour">here</a>.
  `,
  image: require("./image.jpeg")
};
