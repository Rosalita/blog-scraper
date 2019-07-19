So in the last couple of weeks a newly discovered computer security exploit was found - hooray! Something we thought that was safe, trusted, tried and tested over a very long period of time has turned out to be flawed. It's in the media, the sky is falling and people that use the internet are scared about things they do not understand! Customers are frantically emailing companies to ask if they are safe, and how safe safe actually is. An exploit how dreadful, hearts were bleeding last time, what is this nonsense? 


Well it's been named Logjam and it's pretty interesting as it exploits a cryptographic method that was patented back in 1977.



The cryptographic method exploited by Logjam is called a Diffie-Hellman key exchange. So why does anyone actually care and why does cryptography stuff matter?  


Let’s imagine two people that want to share a secret, but they are only able to talk to each other in public where anyone can hear everything they say to each other. That would be a pretty annoying situation, especially as sometimes you really want to share secrets with your best friends without anyone else listening in (I know I do). 


So cryptography solves the problem of sharing secrets in public. The simplest explanation of how a Diffie-Hellman key exchange works is to say it is like mixing various colours of paint together. The trick is that it’s easy to mix two kinds of paint together to make a third colour, and it’s very hard to unmix a paint mixture to establish which two colours made that particular shade.


If these two secretive people wanted to share a secret colour they can do this using a selection of colours of paint. They can both agree in public to start with the same colour like yellow (the public key) then secretly pick a second colour that no one else knows (a private key) which will help them secretly share a new secret colour with each other. 



So let’s say one person's secret colour is red and the other is blue. Both secret colours get mixed with the public yellow colour (to make orange and green respectively). One person then gives the orange paint to the other and receives green paint. The clever bit is now when they add their own secret colour again, mixing red into green and blue into orange, and the end result is they are both left with the same horrible shade of dirty brown. Possibly not the most aesthetically pleasing colour, but no matter how yucky it looks they now both have the same colour and more importantly, nobody else knows what that final colour is. 





[![image](http://2.bp.blogspot.com/-uoKqyIjVHEg/VWydlJFA-eI/AAAAAAAABWw/saaodH8Dc9w/s400/diffiehelman.jpg)](http://2.bp.blogspot.com/-uoKqyIjVHEg/VWydlJFA-eI/AAAAAAAABWw/saaodH8Dc9w/s1600/diffiehelman.jpg)





Colourful diagram shown above because colourful diagrams always help.


Now imagine that instead of mixing colours, a Diffie-Hellman key exchange mixes numbers together instead using hard sums. 


So what happened recently was that some people discovered a way to swap what would be the equivalent of the vibrant paint palettes used by this method for crappier paint palettes. People thought they were picking their secret paint from large palettes containing lots and lots of colours, when unfortunately an attacker had switched their palettes with smaller palettes containing a smaller number of colours. And we all know if you only have a choice of red, yellow and blue it’s much easier to work out that the secret mixed colour will be a nasty shade of brown. 


The logic of mixing the colours was and still appears to be sound, just no-one imagined until recently there would be a way to switch the palettes around and limit the number of private colour choices. As testers we must always strive to imagine the unimaginable, this is one of the reasons why testing is much harder than it appears to be at face value. There may be right answers and wrong answers, but there are also unknown questions which have yet to be answered. Don't worry though, unlike a mythical bug (Rowhammer), Logjam is really easy to patch and most people won't have to do anything more than upgrade their web browser to the latest version to make it go away forever.


 