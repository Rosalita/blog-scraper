Given that games and software testing are two of my favourite things it was only a matter of time before I stumbled upon a game that claims to "test your testing skills". I know what you're thinking, who would even try make a game about testing a game? I'm really not making this up, it exists. Here is the [link](http://www.rapitasystems.com/blog/tetris_coverage_challenge#download-link) if you don't believe me. 




Software that tells someone how good a job they are doing of testing the software - this was too good to miss. I had to download it and take a look.




The way testing Tetris works is that instead of playing for score, the person testing plays for code coverage. Think of functions and statements in the code like Xbox achievements or PlayStation trophies. 100% coverage and the game is won! 





[![image](http://4.bp.blogspot.com/-GmbTdouwNkY/VFYnW6XF62I/AAAAAAAABJM/955e7UUZM8g/s320/tetris.jpg)](http://4.bp.blogspot.com/-GmbTdouwNkY/VFYnW6XF62I/AAAAAAAABJM/955e7UUZM8g/s1600/tetris.jpg)





This screenshot was taken approximately 5 seconds after I rage quit because I suddenly realised what this "game" was doing and why that was bad. The test how well you can test Tetris game is in fact a really good example of how not to be fooled by metrics when testing software. The company which made it sell products that report code coverage metrics. They really want us to believe that test coverage = quality. Bottom line, if something sounds too good to be true, it usual is. 




If we were to say, scale this Tetris example up to a much larger more complex game hmmm maybe something like Skyrim. Yes, Skyrim is a good example. If we played a similar test coverage game with Skyrim and tried to gain 100% coverage how long would it take? a year? two years? ten years? a hundred years? Personally, I didn't work on Skyrim so I would only be able to guess at how it was tested. But I am pretty certain that 100% coverage would have taken too long and cost far too much money to achieve. 100% coverage would have bankrupted Bethesda Game Studios.




But Skyrim was still released, and was widely praised as being a fantastic game. 100% test coverage wasn't critical to the success of the project. If the people playing the game found some obscure bugs in an obscure part of the code (which many did) its wasn't the end of the world. They would just release a patch and the bugs would be gone (hopefully).  




But there is still one more extremely important lesson that software testers can learn take from the sheer ridiculousness of the "test how you test Tetris" example. I'm going to call out to the elephant in the room by saying "Test coverage simply does not equal quality". Please don't ever let anyone trick you into believing it does. This is because test coverage only measures how many lines of code were allowed to run. If a computer could then tell a human if each of those lines of code was correctly implemented or not, every software tester in the world would be out of a job. The bottom line is code with 100% test coverage can still have bugs. 



