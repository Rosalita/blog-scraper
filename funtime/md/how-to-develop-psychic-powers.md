Writing good requirements for software development might seem like an easy task on the surface, but it's actually much harder than many people imagine. The two main challenges that arise when writing requirements are firstly, requirements can change frequently. Secondly, even if you manage to capture a requirement before it changes, it's really easy for someone to [completely misunderstand or misinterpret it](http://www.sandraandwoo.com/2012/11/19/0430-software-engineering-now-with-cats/). Good requirements are absolutely crystal clear with no room for interpretation whatsoever. 



[![image](http://1.bp.blogspot.com/-qouehhp99CA/VdtzIKa2N7I/AAAAAAAABZM/mEtIsCzkffk/s320/psychic.jpg)](http://1.bp.blogspot.com/-qouehhp99CA/VdtzIKa2N7I/AAAAAAAABZM/mEtIsCzkffk/s1600/psychic.jpg)



So what happens when bad requirements happen to good testers? Unfortunately testing software which has poorly documented requirements is far more common than it should be. It's hard to describe the internal thought processes that take place, but its kind of similar to invoking psychic powers. You have to know detailed information about all the things you possess no knowledge of.


Lets imagine the very worst case scenario, you have been asked to test something that has absolutely no written requirements. None, nada, rien, nothing, absolutely no documentation whatsoever.



Warning! This kind of testing is pretty risky, before invoking magic psychic testing powers you should always inform a responsible adult (usually a Project Manager or the Test Manager - should you be lucky enough to have one) about the lack of requirements and associated risks. 


Some of the most common risks encountered will be:


* High priority bugs will be found late. This is because by the time the person doing the testing gains decent knowledge of how the software is actually supposed to work, time will have passed and release date will be closer.  


* The number of 'as designed', 'working as intended' or 'not a bug' defects will significantly increase as testers start making educated guesses as to what might be a bug. 


* Product knowledge will probably only exist inside the heads of 1 or 2 knowledgeable people. The workload of these people will increase as testers try to extract this information from them. It's very rare for knowledge holders to be available to answer questions all the time.


* Test automation will either grind to a halt or happen very late. How can you write automated regression tests if you don't know how the product is supposed to work? The simple answer is you can't. 


So once you have told the responsible adult in charge about the risks of testing with no requirements they may say something along the lines of ''We can't write requirements because no-one knows how it works.' It could be a legacy product you're being asked to test. It could be that the person that created it left the company without writing any kind of documentation. You may even be told 'We simply don't have time to write any requirements'.



[![image](http://4.bp.blogspot.com/-5vTNc4tPXbE/Vdtz5-CRBvI/AAAAAAAABZU/LkkIGsal7vA/s320/psychic2.jpg)](http://4.bp.blogspot.com/-5vTNc4tPXbE/Vdtz5-CRBvI/AAAAAAAABZU/LkkIGsal7vA/s1600/psychic2.jpg)



What happens now? Don't panic, I'm going to try guide you through the most efficient pain free way to test the unknown. The following approaches can help maximise testing efforts while also giving the illusion that you have developed some kind of psychic testing ability.


At the most basic level any testing carried out on a requirement-less project will fall into two categories.


Category 1 - Obvious things - I'm certain if I do this, the software should do that.


Category 2 - Mystery things - I have no idea what the software is doing, why it is doing it or even if it should be doing it at all. 


An example of a category 1 obvious thing would be a text input box that says 'email address' with a button below it that says 'subscribe to newsletter'. A fairly safe assumption would be that entering a valid email address and clicking the button will subscribe the email address to a newsletter.


A category 2 mystery thing might be an unlabelled text input box with a button below it that says 'Start'. What is being started? What should happen when it starts? How do I know if it actually started? What should be typed into the input box?


A good tester will explore the software and be able to draw from a number of sources to guess the errors. All the points listed below have all worked for me in the past when I have been expected to test unknown entities.  


* Try to test important and critical features first however without requirement to work from, it may not be immediately obvious which features are critically important. So start with the obvious functionality which is basically everything that falls into category 1.


* Break the software down into smaller areas or sections. Keep track of all the obvious things that were tested in each of these areas and what the results were. This information can be used as the starting point to form regression tests.


* While you are breaking the software down into smaller component pieces and testing all the obvious things, questions will come into your mind about the features that fall into mystery category 2. Compile a list of questions about the mysterious features that fall into category 2.


* All the time you are doing this, rely on your instincts! If something feels like a bug because it's acting in an unexpected way then it's highly likely it's a bug - even if that bug might turn out to just be a poor design choice. Anything that detracts from the overall quality of the software should be considered a bug.


* Seek answers to the mystery questions. How does the functionality compare to the previous version or to a competitors product? These insights can give valuable clues as to whether or not something is working correctly. Learn as much as possible about the product's functionality from reliable sources.


* Ask developers how they expect the software will behave. If you don't have any requirements to test against, it's likely your developers didn't have requirements to develop against either but they should at least be able to tell you what kind of functionality they added.


* Always keep notes while exploring and learning about the software. Document unguessable things once you discover how they are supposed to work. Trust me, it will save a great deal of time later when you have to revisit complex areas and remember what's going on. There is also bonus value in having notes should a new tester join your project and you need to get them up to speed quickly or if you ever find yourself in the situation where you have to hand over your testing work to someone else.


* If in time doubts arise as to whether or not to log a bug, just log the bug. Once it is entered into a defect tracking system people are usually very fast to point out false positives and it only takes a moment to close them down. 


* Try to confirm your test results with anyone that already holds expert knowledge of the product. Remember all your test results are still just assumptions until they are confirmed or denied.


Whatever you do don't give up or get disheartened. While lack of well documented requirements and user stories certainly increases the difficulty level of testing, it certainly doesn't make testing completely impossible. Always do the best you can with the tools and information you have available to you.