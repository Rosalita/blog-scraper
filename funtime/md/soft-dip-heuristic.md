Just as different people can possess different political beliefs and not everyone believes the same thing, I think the same can be said with software testing. In the world of testing there isn't a one size fits all 'right answer', it doesn't exist. Lots of people have lots of different ideas and some of these ideas can conflict with each other.  The whole manual vs. automation argument is a good example of this. Some people think that automated testing is a silver bullet that will eliminate all bugs. Some people believe that test automation is so expensive in terms of time and effort in relation to the value it returns that it should be used sparingly.  


When I think about where my testing beliefs fit into the testing community around me, the principles of [context driven testing](http://context-driven-testing.com/) resonate with my personal experience, like music in my ears. Testing is hard, I know this, I have experienced this first hand and as such I know that there are no absolute best practices. What might be a good approach for one problem could be totally impractical or impossible for a different problem. But I also know there are ways we can make life easier for ourselves.


James Bach seems to really like making up nonsensical acronyms to try remember good ways to test things. Some examples of nonsensical testing acronyms created by James Bach would be  



HICCUPPSF
 =  
H
istory, 
I
mage, 
C
omparable Product, 
C
laims, 
U
ser 
E
xpectations, 
P
roduct, 
P
urpose, 
S
tandards and Statutes, 
F
amiliar Problems 



CRUSSPIC
 =  
C
apability, 
R
eliability, 
U
sability, 
S
ecurity, 
S
calability, 
P
erformance, 
I
nstallability, 
C
ompatibility 



CIDTESTD
  =  
C
ustomers, 
I
nformation, 
D
eveloper Relations, 
T
eam, 
E
quipment & Tools, 
S
chedule, 
T
est Items, 
D
eliverables 



DUFFSSCRA
  = 
D
omain, 
U
ser, 
F
unction, 
F
low, 
S
tress, 
S
cenario, 
C
laims, 
R
isk, 
A
utomatic


If you don't believe me, paste any of those seemingly random strings of capital letters into Google and see what comes back :)


My favourite of all these 'Bachist' acronyms is the SFDIPOT one, because even though it at first glance it sounds like utter bollocks it's the one that has proven the most useful to me and as a believer in context driven testing, I only care about practices that are actually good in context.  I still thought the way he had arranged the letters made it sound like bollocks, so I rearranged it in my head so I could remember it easier. After all this is about what works for me, not what works for Mr Bach.


You say potato, I say potato. You say SFDIPOT, I say SOFTDIP. Soft dip is nice, especially at parties with an assortment of breadsticks and savoury nibbles.  What is SOFTDIP? 




[![image](http://2.bp.blogspot.com/-nXn8o9Nbg_E/VYxGv-ks9II/AAAAAAAABXc/Iq8SPBRQB0s/s320/breadstick.jpg)](http://2.bp.blogspot.com/-nXn8o9Nbg_E/VYxGv-ks9II/AAAAAAAABXc/Iq8SPBRQB0s/s1600/breadstick.jpg)




SOFTDIP
  = 
S
tructure, 
O
peration, 
F
unction, 
T
ime, 
D
ata, 
I
nterface, 
P
latform


Each of the words asks questions about the thing being tested. I find that asking these questions help me to imagine how I am going to test, identify important scenarios before I begin testing and make sure I don't overlook anything really important.


The Softdip questions I ask myself are:



Structure
 - what is it made of? how was it built? it is modular can I test it module by module? Does it use memcache? Does it use AJAX?



Operation
 - how will it be used? what will it be used for? Has anyone actually given any consideration as to why we need this? are there things that some users are more likely to do than others?



Functionality
 - What are its functions? what kind of error handling does it have? What are the individual things that it does? Does it do anything that is invisible to the user? 



Time
 - Is it sensitive to timing or sequencing? Multiple clicking triggers multiple events? How is it affected by the passage of time? Does it interact with things that have start dates / end dates or expiry dates?



Data
 - What kind of inputs does it process? What does its output look like? What kinds of modes or states can it be in? What happens if it interacts with good data? What happens if it interacts with bad data? Is it creating, reading, updating and deleting data correctly? 



Interface
 - How does the user interact with it? If it receives input from users, is it possible to inject HTML/SQL? What happens if the user uses the interface in an unexpected way? 



Platform
 - how does it interact with its environment? Does it need to be configured in a special way? Does it depend on third party software, third party APIs, does it use things like Amazon s3 buckets? What is it dependent on? What is dependent on it? Does it have APIs? How does it interact with other things?


Yeah, I know what you're thinking just another mnemonic but let me give you an example and show you how this works for me, because if it works for me who knows, it might work for you. 

Once upon a time I was given some work to test, a new feature which when used correctly will delete a customer from a piece of software. Sounds simple doesn't it on the surface. The specification for the feature was just that, we need a way to delete a customer from the software. How would you test that? What would you consider? I know from experience that deleting things is a high risk operation so I wanted to be certain beyond doubt that this new feature would be safe and have no unexpected consequences or side effects. So I worked through SOFTDIP and this is what I came up with.



Structure
 - I have been provided with a diagram that shows where all client data is stored in the software's database. It seems that information about a customer can be present in up to 27 different database tables. When I test this I will need to make sure that no traces of unwanted data are left behind. Maybe I can write some SQL to help me efficiently check these 27 tables.



Operation
 -  The software is legacy and has so much data in it currently that it costs a lot of money to run. As customers migrate to our newer product, their old data is being abandoned in this legacy software and we are still paying for it to be there. This is why the feature is needed. Only admin users will use this feature to remove data, customers and end users must not be able to access this feature. I am going to need to test that non-admin users are definitely unable to access this feature.



Functionality
 -  It deletes existing user data. It must only delete the data it is told to delete and leave all other data intact. It must delete all traces of the data it is told to delete.



Time
 - What happens if deletion is triggered multiple times? What happens if removing large amounts of data takes too long and the server times out? What happens if the act of removing data is interrupted? What happens if data is partially deleted and an admin user attempts to delete it again? 



Data
 - The software can be in the state of not deleting anything or deleting stuff. Live data is HUGE compared to the amount of data in the staging environment. I will need to prove that this feature functions correctly running on live data. This will most likely involve paying for an Amazon RDS to hold a copy of live data. I need to make sure I know exactly what I'm going to test and how before I request this RDS to minimise costs. It's also possibly going to be a good idea to take a database snapshot taken of this RDS before testing starts so can easily restore back to the snapshot if/when tests fail or need to be re-run.



Interface
 - I have been told there will be 'something' at URL ending /admin which will allow the admin user to delete account data. I need to make sure that customers are not able to access this URL directly and that only admin users are able to initiate the deletion process. I'm also going to have to make sure that even though this interface wont be seen by customers that it is fit for purpose. Consideration should still be given to things like asking the user for confirmation before any kind of deletion starts.



Platform
 -  This software does make some requests to third party software to retrieve data however if customers are deleted then those requests won't happen as the software won't know what to ask about. I need to prove that this statement is true. There is a second piece of software that asks the piece of software in test for data, what happens if it asks about a client that has been deleted? I'm going to have to make sure that I also test this scenario. 


Asking these questions starts to forms the basis for my test plan

* attempt to access feature as non-admin user, verify non-admin unable to access feature

* Make sure user is asked to confirm before delete operation starts

* attempt to start the delete operation multiple times for the same customer

* attempt to start the delete operation multiple times for multiple customers

* ensure feature works correctly when using live data

* ensure after running delete operation that all data has been successfully removed

* carry out a regression test to ensure existing functionality is maintained

* test the delete operation for a customer which has a really large amount of data 

* verify software no longer makes requests to third party software for deleted customers 

* verify that other software which makes requests to the software in test still functions 


So why bother going to all this trouble, all of this preparation before starting testing? Well I'm always happier when I run code changes against a decent test plan. It makes me feel reassured when the time comes to release them into the wilds of the production environment. Every day, a lot of people depend on the software I test and I feel a strong responsibility to them to do the best job that I possibly can. Good testers care deeply about creating good software.