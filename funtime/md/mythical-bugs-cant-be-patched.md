When you work in software testing, every now and then, you get to hear other people's stories about bugs. Most of these stories will be fairly mundane. Something along the lines of like "Yeah, I clicked the button and nothing happened". But there will be other times, once you have been working in software testing for a while, when you may get to hear a story about a legendary bug. Legendary bug stories tend to something like "Yeah, and then if you paste that in to the text box and hold down the shift key, it sends an email to 193,248,2489 customers thanking them for ordering Nickleback's latest album". 



I love legendary bug stories. They can serve as mild amusement, shining examples of things that should or should not be done, cautionary tales of woe or even be so far outside of the box that they change the way a tester will think about certain problems or situations. I think all testers must love a good bug story, the ones I go drinking with certainly do :) 


The legendary bug story from my last games testing team went something like this. Once upon a time there was a racing game that was about to be released. While testing this racing game, one of the AI controlled cars fell through the track. Replays of the race had to be watched back and scrutinised in microscopic detail to try ascertain which car had fallen though the track and more importantly where it was on the track when it fell through. All the cars had the names of the drivers written just above the wind-shields. While watching the replay footage back, one of the testers noticed a spelling mistake. A driver called 'Bayne' had 'Banye' written on his car. This spelling mistake had been in the game for a long time. The spelling mistake had been missed by everyone, at every level of development and was also in a whole load of promotional screen-shots and marketing material for the game! This legendary bug would possibly fall into the cautionary tales of woe category. The fact that test caught it very late in the day and saved the company from significant embarrassment pushed the story of 'Bayne' into legendary bug status. Seriously, I'm surprised noone on the team submitted it to [The Trenches](http://trenchescomic.com/). 


Outside the world of games however, legendary bugs can sometimes be utterly mythical. At Google they have an all-star team of security testers dubbed 'Project Zero'. These people actively hunt out vulnerabilities with the aim to find the flaws before the bad guys find them so they can be fixed. Well, Project Zero found a new bug last week. Not just any bug, a mythical a hardware bug!


The story goes like this. Computers use memory to remember things. There is a type of memory called Dynamic random-access memory (abbreviated to DRAM). DRAM works by storing every bit of data in a separate capacitor. The capacitor can be either charged or discharged and these two states represent 0 or 1. Google's testers found a way to change 0's to 1's or 1's to 0's without accessing them. They found that if you pick two memory locations either side of third memory location, and bombard these two 'aggressor' locations with requests, the third 'victim' location will just flip from 0 to 1 on its own. 


They are calling the exploit Rowhammer and you can read the [Project Zero blog post here.](http://googleprojectzero.blogspot.in/2015/03/exploiting-dram-rowhammer-bug-to-gain.html) The worst thing about this bug is that it is physical in nature. It can't be patched.


There is currently a [test for Rowhammer](https://github.com/google/rowhammer-test) on github.com although in the warnings it does say "Be careful not to run this test on machines that contain important data." So you possibly won't want to try this on your home PC. At least knowledge of this issue is in the public domain now. Knowing about the Rowhammer exploit exists possibly makes it slightly less terrifying. It certainly will be interesting to see if and how anyone takes advantage of it.



[![image](http://2.bp.blogspot.com/-H6YXJy8Sis0/VQdSoSDc-NI/AAAAAAAABU0/E0hUX_-dvCk/s400/bug.png)](http://2.bp.blogspot.com/-H6YXJy8Sis0/VQdSoSDc-NI/AAAAAAAABU0/E0hUX_-dvCk/s1600/bug.png)
