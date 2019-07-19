It's been a roller-coaster of a month since my last blog post. In the last four weeks I have successfully managed to change job and learn JavaScript! I started on JavaScript the same way as Python by completing the free [codecademy](www.codecademy.com) course. If you test things and you want to learn basic programming you should definitely give it a try.





[![image](http://3.bp.blogspot.com/-xTC_hA4OiJw/VdUFNbraaRI/AAAAAAAABYU/KIuIu4ZtVSk/s320/JS.png)](http://3.bp.blogspot.com/-xTC_hA4OiJw/VdUFNbraaRI/AAAAAAAABYU/KIuIu4ZtVSk/s1600/JS.png)





Some initial observations made while learning JavaScript:


1) The learning process was much faster than last time. Knowing a first language definitely helps with learning a second. My first [Fizzbuzz](http://c2.com/cgi/wiki?FizzBuzzTest) in Python took 30 minutes, but my first [Fizzbuzz](http://c2.com/cgi/wiki?FizzBuzzTest) in JavaScript took 3 minutes.


2) White space is not an enemy in JavaScript land. Viva the curly bracket!


3) Forgetting semicolons isn't nearly as bad as I thought it would be. 



I've also learned absolutely loads of things about test automation with JavaScript in the last couple of weeks, which is the main reason for this blog post (hooray!). 


One of the first things I did was install  [Node.js](https://nodejs.org/) which comes with a truly awesome package manager called [npm](https://www.npmjs.com/).  The package manager made it really easy to try out all of these testing frameworks. Beware if you're on Windows 10 however, some tweaking was required to get it working correctly (Stack Overflow is your friend).


I discovered that there are many different testing frameworks available for writing tests with JavaScript. Actually it's not just testing frameworks, there are many, many JavaScript frameworks in general. Far too many of them. There is a joke among developers that  [a new JavaScript framework is born every sixteen minutes!](http://www.allenpike.com/2015/javascript-framework-fatigue/)



Testing frameworks I encountered and explored were:


* [Jasmine](http://jasmine.github.io/2.3/introduction.html)


* [Mocha] (https://mochajs.org/) 


* [Chai](http://chaijs.com/) 


* [Cucumber.js](https://cucumber.io/docs/reference/javascript)


* [Selenium WebDriver JS](https://code.google.com/p/selenium/wiki/WebDriverJs) 


* [Nightwatch.js](http://nightwatchjs.org/) 


* [Protractor.js](https://angular.github.io/protractor/#/)  




Some of these frameworks are specifically for unit testing, some are for end to end testing. Some depend on each other, some are agnostic and framework free.


I drew a little ASCII diagram to try visualise them. Each framework is listed left to right in a box with either (u) for unit testing or (e2e) for end to end testing. Each framework box has everything it uses on listed underneath it. 



[![image](http://2.bp.blogspot.com/-EsFxw2GMBNk/VdVlJEnfeHI/AAAAAAAABZA/Phrkh2HmktY/s1600/ascii2.png)](http://2.bp.blogspot.com/-EsFxw2GMBNk/VdVlJEnfeHI/AAAAAAAABZA/Phrkh2HmktY/s1600/ascii2.png)



These test frameworks increase in complexity from left to right. Jasmine stand alone is a simple unit test framework that just requires JavaScript. Protractor is a more complex end to end test framework that requires either Jasmine (or Mocha and Chai) (or Cucumber) and uses both WebDriver and Node.js 



I had a play around with [Jasmine stand alone](https://github.com/jasmine/jasmine/releases) but as this is a unit test framework, I found I had to actually write some Javascript code before I had anything to run my tests against.  Unit tests are usually written by the developers that are developing the application. As a Test Engineer, the tests I need to write are a mixture of both acceptance tests, integration tests and end to end tests. 



* [Acceptance test](https://en.wikipedia.org/wiki/Acceptance_testing) - Determines if a specification (also known as a user story) has been met. 


* [Integration test](https://en.wikipedia.org/wiki/Integration_testing) - Determines if a number of smaller units or modules work together. 


* [End to end test](http://www.techopedia.com/definition/7035/end-to-end-test) - Follows the flow through the application from the start to the end through all the integrated components and modules. 


I looked at Protractor next. Protractor is a testing framework which has been around for a couple of years. I saw that the tests were formatted in a [BDD] (https://en.wikipedia.org/wiki/Behavior-driven_development)(Behaviour Driven Development, not Beer Driven Development) style. 


The syntax protractor uses is based on expect along the lines of 'expect something to equal something else' rather than the more familiar verify/assert statements I encountered when I was writing Selenium WebDriver tests in Python. Protractor's main strength is that it was created specifically to test AngularJS applications. It supports element location strategies for Angular specific elements. If you need to test anything created in AngularJS, Protractor is the King of the Hill.



I then moved on to looking at Nightwatch which felt closer in syntax to the Selenium WebDriver tests I had previously written. Nightwatch is newer than Protractor making it its first appearance on GitHub in February 2014. I found a  [good tutorial for getting started with Nightwatch which also has a demo on GitHub] (http://matthewroach.me/ui-testing-with-nightwatch-js/).


I after a bit of playing around with it, I decided I was going to re-write my Python Pinterest test in JavaScript Nightwatch.


I went through all the Nightwatch asserts and commands and tried to include as many of them as possible in the sample test I wrote.


It was very reassuring to see first-hand that JavaScript and Nightwatch are capable of carrying out all of the tasks possible with Python and Selenium WebDriver.


Anyway here is the [test example I wrote with JavaScript and Nightwatch](https://github.com/Rosalita/Selenium/blob/master/JavaScript%20Examples/Nightwatch/tests/pinterest.js). One of the main advantages I found of writing within a testing framework was that creating the tests was actually much faster. The amount of text I had to physically type in was less than if I hadn't been using a test framework. Also, instead of faffing around with variable assignments a lot of the nitty gritty of what was going on in the background was actually hidden away from me allowing me to just focus on writing the test. 


 