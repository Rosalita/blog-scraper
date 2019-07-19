I saw [this article](http://www.birminghammail.co.uk/news/midlands-news/23-billion-ticket-birmingham-airport-7836398) recently about how a website owned by a company called eDreams nearly charged a lady £23 billion for a return flight.




Closer inspection revealed that it wasn't the flight that caused the problem, it was actually the return baggage check-in cost.





[![image](http://2.bp.blogspot.com/-ofGkJ6733Hg/VGpvURjwzxI/AAAAAAAABQI/JCmpZo0c8X0/s400/air-ticket-main.jpg)](http://2.bp.blogspot.com/-ofGkJ6733Hg/VGpvURjwzxI/AAAAAAAABQI/JCmpZo0c8X0/s1600/air-ticket-main.jpg)





How could such a massive error make it unspotted onto the eDreams website? 


Maybe they didn't have any automated tests that could verify the return baggage cost was far to large. 

Maybe the bug was an edge case that only occurred under very specific circumstances, which this customer accidently stumbled upon. 

Maybe the testers were outsourced and just didn't care. 

Maybe the testers didn't have control of their staging environment and were trying to test a moving target as their developers kept updating it every 20 minutes. 

Maybe they didn't have a staging environment at all.

Maybe the testers were not given enough time to test. 

Maybe their testing department found the bug, but didn't have the power to stop the build from being released.

Maybe the data the testers tested with was different to live data.

Maybe the customer followed a different path through the software to the testers.

Maybe none of the testers were looking at the return baggage check in cost, because that area had not changed recently.

Maybe the website simply wasn't tested at all.


There are potentially hundreds of reasons this bug could have been missed.


What damage has the bug caused? I would say massive damage to their company reputation. The £23 billion ticket made headline news! As a customer I would be wary of buying anything from their website. As a software tester, this does not sound like a company I would EVER want to work for.


About a month after the appearance of this bug on their website, [British Airways and Iberia decided to withdraw their fares](http://www.reuters.com/article/2014/10/24/edreams-odigeo-suspension-idUSL6N0SJ47Z20141024) from three of their websites. This, in turn, wiped 59% off the e-dreams share value. How long now until this company folds?


When software goes bad, it can go very, very, very bad. Some people don't realise just how bad this can be. It's doubtful that there will be a happy ending for eDreams. Their story is a true software testing nightmare. 