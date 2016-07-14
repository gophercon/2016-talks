# CONTRIBUTING: The Talk!

Michael Matloob 11 July 2016

Interested in contributing to the Go Programming Language open source project,
but don’t know where to start? Come learn how you can get involved and find
something to work on.

If you’re already contributing, come learn how to more effectively review CLs
and help make the contribution process smoother for others.

--------------------------------------------------------------------------------

Hi!

# Contributing, The Talk!

![title](slides/title.png)

Welcome to "CONTRIBUTING: The Talk"! My goal for this talk is to help those of
you who want to contribute to the Go Programming Language Open Source project to
get involved and start contributing. We've got some stuff for veteran
contributors too near the end.

# matloob's story

![matloob's story](slides/matloobs_story.png)

This is the talk I wish I'd seen when I started contributing to the Go project.

When I started contributing, I didn't how to start. I took on a linker bug that
I didn't know how to fix, and I got really intimidated by the comments on the
changes that I sent out. So intimidated that I took a break for more than a
year.

The second time around, I was better equipped. I found more a suitable task to
work on, a tool to move go packages to a new path in a source tree. And as I
made changes, I eventually figured out some of the rules and conventions in
contributing. I realized that it was okay to receive critical comments on my
changes -- the reviewers weren't out to get me, the comments were just their way
of teaching me how to contribute. And I've grown comfortable with some parts of
the code base, such as our tools and the regular expressions package.

This talk is an attempt to help you get started in becoming a contributor.
Hopefully it can save you some time on your path...
        
# Don't be (too) intimidated!

![](slides/dont_be_too_intimidated.png)

Contributing to Go may seem intimidating at times. And it is intimidating! It's
important code that a lot of people depend on. Even the core developers on the
project are mindful of making only the right changes and often have their
proposals rejected. The Go project is not an experimental project anymore --
it's serious business for a lot of people.

But you shouldn't be too intimidated! There are lots of ways to get involved for
every level of experience, and I'll show you a couple of options in this talk.

# It's all about the people! Be nice!

![](slides/be_nice.png)

We shouldn't forget that as contributors, we're not just working with _code_,
but also with people. It goes without saying that contributors follow the
Code of Conduct, but there's more than just following the rules -- we should be
nice to each other!

There's no non-verbal context when communicating on our mailing lists, so
consider how your words might seem to someone who can't see you. There are a lot
of times when we'll have to say no to a proposal or a change. If you're already
a contributor, and it's your responsibility to say no to your fellow
contributor, you should be careful to communicate that it's not personal, and
you still welcome them in our community.

Conversely, when you're getting feedback and constructive criticism about your
work from other members in the community, it's easy to read it in a negative
subtext. Sometimes I read comments from people that say to "change X" and I
might still read it as "change X, stupid!" but that's not what they said. They
just want Go to be the best it can be, just like I do. So don't get scared if
you see your changelists bleeding with red ink...

When you're a beginner, those early rounds of review are part of the learning
process. As you anticipate our patterns, edits to later CLs become lighter. And
the whole time, our code's getting better too!

With all this said, it's good to keep in mind that everyone involved in Go is
very, very busy, and since we're human we sometimes don't take enough time to
communicate carefully.

When that happens, give the reviewer the benefit of the doubt. If you're upset
by it, reach out to them off list to clear the air.

![](slides/win_friends.png)

On a personal note, I've found the book "How to win friends and influence
people" by Dale Carnegie
to be really useful for this 'communication' thing. It's full of a lot
of common-sense that's really easy to forget sometimes...

# Overview

![](slides/overview.png)

Alright. The rest of the talk is structured as follows. First, I'll go over the
components in the Go project, and the tools and resources you'll be using as a
contributor.

I'll then talk about some ways you can get started as a contributor, and give a
demo of making a code contribution.

I'll finish with tips and advice for more complex contributors.

# The go project

![](slides/the_go_project.png)

Let's take a look at the components of the go project.

![](slides/the_go_core.png)

The core repository contains the compiler and runtime, standard library and
docs for Go. It's basically the stuff that's distributed in the standard
Go installation.

![](slides/the_subrepos.png)

The subrepos contain packages and other tools that have been developed as part
of the project. For instance, 'tools' contains tools such as guru and goimports,
and 'mobile' contains support for developing Go apps that run on iOS or android.

These packages have imports paths that start with `golang.org/x/tools`

# The tools

![](slides/the_tools.png)

Now, let's take a look at the tools.

# golang.org

![golang.org](slides/golang_org.png)

golang.org has a number of resources for contributors. Of particular interest to
contributors we have...

"The Go Project" page, which links to a number of resources about the project,
especially

![contributing guidelines](slides/golang_org_contribute.png) 

the "Contribution Guidelines". The contribution guidelines contain a lot of
the contribution instructions in this talk. You should definitely take a look
if you're interested in contributing!

Don't worry about everything up here, I'll go over all of this later in the talk.

# Github

![github](slides/github.png)

We use Github primarily for its issue tracker. It also hosts our wiki.

This is what the go github page looks like.

![](slides/github_mirror.png)

We don't accept pull requests _through github_ we use
a much better system called Gerrit! More on Gerrit later.

It's not our primary Git repository, it's just a mirror.

# go.googlesource.com

![go.googlesource.com](slides/go_googlesource_com.png)

go.googlesource.com is the official source for the Go source. Both the go
core repository and the subrepositories are hosted here. You can browse the git
log and branches and download the code here too.

![](slides/golang_org_x_magic.png)

The subrepos, which have import paths that start with golang.org/x/
are also stored on googlesource. We use a feature of go get to redirect
golang.org/x/ to the appropriate git repos.

# golang.org/cl

![golang.org/cl](slides/gerrit.png)

And we use a different codereview tool also. It's called Gerrit.

This is where new code contributors add to the go project goes through a process of
peer review.

For now let's take a look at what it looks like

![](slides/gerrit_change.png)

If we click on one of the entries, we get the page for that entry.

Gerrit allows us to review code by posting comments
and assigning scores. 

![](slides/gerrit_discussion.png)

Those show up on the email discussion thread for the code review.
This is what it looks like on gerrit.

We can add those by clicking on the 'reply' button.

Gerrit has a notion of scores...

![](slides/gerrit_zero.png)

A zero score is neutral.

By the way, you may not see the +2 and -2 and the trybot stuff as a beginning
contributor. They'll appear once you're a committer.

![](slides/gerrit_minus_two.png)

A minus two means that the reviewer thinks there's a significant problem with the change
that needs to be fixed. It blocks a change from being submitted  until
it's retracted.


![](slides/gerrit_minus_one.png)

We don't really use minus ones much.

![](slides/gerrit_plus_one.png)

A plus one means the reviewer likes the change, but thinks it needs approval from another
contributor before it's submitted.

![](slides/gerrit_plus_two.png)

A plus two is that approval. It means that the reviewer thinks the change is ready to be submitted. Every change
needs to get a +2 before it gets submitted.

This particular change already has a plus two, as you can see on the left hand side
of the screen.

We'll talk more about what codereview later in the talk.


# End to end example

![subtests example](slides/example_subtests.png)

Let's go through an end-to-end example of making a change to Go.

Our example will be the new subtests feature coming in 1.7.

Before, each named test case or benchmark would have to be written as its own function.
Here, in our regular expressions benchmark there would need be one benchmark for each combination of input string length and
regular expression.
This feature allows us write multiple benchmarks in the same function.
It's pretty nifty.

![end to end issue](slides/example_subtests_issue.png)

This feature was contributed by Marcel, one of our contributors,
who has the github handle mpvl.

Let's take a look at the steps Marcel went through to add this feature.

First, Marcel makes a proposal issue.

This is the Github issue page where the proposal was made.

The proposal issue contains a description of the feature and a rationale for
the feature.

By the way, that golang.org link at the bottom is a convenience link
that redirects to github.

![end_to_end_proposal](slides/example_subtests_proposal.png)

Then, since this is a pretty major feature, we put together a proposal document.

Proposal documents are sent as markdown files and checked into the 'proposal'
subrepo. This is what the checked in proposal looks like as rendered by Github.

This document has sections introducing the problem being solved with the change,
a through description of the change being proposed and how it solves the problem,
an exampination of alternatives, and an implementation strategy.

The discussion of the change continues
on the GutHub issue. Eventually consensus is reached on the issue.
Now Marcel can start writing code for the implementation.

![end_to_end_change](slides/example_subtests_polygerrit.png)

And then, send changes, this feature required a number of changes, this is just
the last one. Notice the +2 approving the change to be submitted.

This is the main gerrit codereview page for this particular change.
We have a special  golang.org/cl link that redirects to the change.

Oh, and I'm showing you the new, shiny gerrit UI, called polygerrit.
You can follow that link on the bottom of this slide to toggle it on.

![end_to_end diff](slides/example_subtests_diff.png)

Here's an example of what codereview looks like. Marcel sent out the change,
Russ wrote a comment (on the comments) and then Marcel updated a new version
of the change responding to the comment.]

# Beginners!

![](slides/beginners.png)

If you're a beginner contributor to the Go project we'd recommend starting by
filing bugs, updating docs, or adding test cases.

## Filing bugs... and investigating them

![](slides/golang_org_issue_new.png)

You probably won't go hunting for bugs when you're a beginner, because bugs
often show up when you least expect it. But if you do encounter a bug, we'd love
a report. And there's no harm in filing duplicates! So even if it seems like we
must have already seen a bug, you can send the report anyway!

To file a bug, go to golang.org/issue/new, which will pop up a form with
prompts for details about your issue.

This slide shows what that form looks like.

And if you see a bug already in the tracker, but don't know how to fix it, you
can still help! Investigating and clarifying bugs is extremely helpful. Creating
a test case or reproduction for a bug, (or minimizing a complex one) is a huge
help in getting it fixed.

## Updating docs

![](slides/beginners_docs.png)

No nit is too small when it comes to our docs. If there's anything that's not
clear, either in godoc comments or the documentation on golang.org, Just make a
change to the go git repository and send it to us.

![golang_org_doc](slides/golang_org_doc.png)

By the way, the documentation on golang.org is contained in the go repository in
the 'doc' directory. You can see golang.org/docs on the left side of this slide.

On the right side of the slide, you can see our wiki.
It's on github, and it's open to edits by any github user. Go to
golang.org/wiki and click the edit button to make changes there...

## Adding test cases

![](slides/beginners_test_cases.png)

Another good way for a beginner to get involved is by contributing test cases.
Or benchmarks -- they're kind of like test cases... I'll demonstrate this...
In a DEMO!!!

![](slides/re2_dfa.png)

So... I've recently been working on a change to the regexp
package. It's a port of the fast (but complicated) RE2 DFA matcher to Go.
But there were three [sic] problems: first-- I didn't have my
work done in time before the freeze (more on the freeze later) and second-- I
didn't have enough of a justification that the complexity was worth it.

    ( the third is that paper russ was talking about... )

One way of making that justification is to demonstrate an improvement on the benchmarks, but
the benchmarks that are currently checked in are pretty easy for the current NFA
matcher... Let's try checking in something a bit harder...

![](slides/re2_fanout.png)

Such as this one from the C++ RE2 regular expressions implementation... Oh look, "NFA
execution will be particularly slow"... just what we need!

Don't worry too much about the C++ code on this page. The important part here is
the string containing the regular expression. We'll be able to copy it straight
into the Go regular expressions benchmarks.

## DEMO TIME!

![](slides/demo_time.png)

By the way, everything I'm about to demonstrate is in all.bash.

Let's add that benchmark to the Go regular expressions benchmark suite. The change I'm going
to make is to add that fanout testcase to the regexp benchmark we saw earlier.

Okay, let's start... I'll assume that you've already installed Go using one of
the installers you can download on "golang.org/dl". On my Mac, that puts the Go
tree in /usr/local/go. Okay, let me set up my shell environment

    export EDITOR="code"
    export GOPATH="$HOME"
    export GOROOT_BOOTSTRAP=/usr/local/go
    PATH="$GOPATH/bin:$PATH"

So I'm setting my EDITOR env variable and I'm also setting my GOPATH to $HOME,
both of which are personal preferences. I'm also setting this special
GOROOT_BOOTSTRAP environment variable. This is only necessary if you want to
build the Go compiler toolchain or standard library in the core repository. The
Go compiler toolchain is built using Go and we need to have an already existing go
installation at GOROOT_BOOTSTRAP to do the building. I'm also adding the my
GOPATH's bin directory to my path. We'll be go getting some tools that will help
us with the contributing process. But first, let's get the code we're going to
hack on.

It's available at the go git repository at go.googlesource.com. Let's git clone
it. I'm just going to put it in my home directory.

    git clone https://go.googlesource.com/go

![](slides/make_bash.png)

Okay, now let's build it. We've got a script called "all.bash" in the src
directory that builds and tests the go tree. It's a good idea to run it after
downloading the tree to make sure everything's in working order. We're going to
run make.bash in this demo because it takes less time.

    pushd go/src
    ./make.bash

        ##### Building Go bootstrap tool.
        cmd/dist
        ...

All right, this is going to take a while! Now let's get those tools!

    go get -u golang.org/x/review/git-codereview

Let's also install some shortcuts in git. These will make it a little easier to
use the codereview tools we just installed...

    $EDITOR .gitconfig

    [user]
        name = "Michael Matloob"
        user = "matloob@golang.org"

    [alias]
        change = codereview change
        gofmt = codereview gofmt
        mail = codereview mail
        pending = codereview pending
        submit = codereview submit
        sync = codereview sync

Great...

Let's check back on our build...

    .....
    cmd/compile/internal/x86
    cmd/compile


    ---
    Installed Go for darwin/amd64 in /Users/matloob/go
    Installed commands in /Users/matloob/go/bin

Ok, now let's make the change.

First I'm going to use `git change` to make a new branch and switch to it...
In our code contribution model, you'll prepare a single commit on top of our
git repository's HEAD that we'll merge in. `git change` assists with that.
    git change fanout

The file we're changing is regexp/exec_test.go

(add the following code)

    {"Fanout", "(?:[\\x{80}-\\x{10FFFF}]?){100}[\\x{80}-\\x{10FFFF}]"}

great!

Now let's test things

So we want to run the regular expressions tests and benchmark. Usually, we use

    go test regexp

to run the regular expressions tests and benchmarks. But the go tool has
embedded within it the path to its source code, so go test regexp is going to
run the tests for our stable go installation

What we want to do here is to use the version of go we're hacking on.
all.bash puts the built go binary in $HOME/go/bin/go. I'll make an alias
to that version of go

    alias dev-go="$HOME/go/bin/go"

Cool. And

    $HOME/go/bin/go test regexp -short
    
will run the regexp tests

Alright. Let's run the new benchmarks

    go test regexp -run None -bench BenchmarkMatch$/Fanout

This command tells the testing tool to not run any tests, and to only
run the fanout benchmarks. It uses the new subbenchmark matching functionality
added by Marcel.

cool!

## Change is ready!

![](slides/change_is_ready.png)

## CL descriptions

![](slides/change_descriptions.png)

We use change descriptions to communicate both to our reviewer and to future
contributors. Our reviewer wants to understand what the change is adding as context
when reviewing the change. A future contributor may want to understand why a
certain change was made.

Our change descriptions have a certain structure you should try to follow.

![](slides/anatomy_change_description.png)

Let's take a look at the structure of a change description:

(by the way, this is what the change looks like on googlesource.com)


![](slides/anatomy_change_description_first_line.png)


The path we're making the change to. This is the subdirectory of src in for
the core repo.

It's `golang.org/x/[the name of the subrepo]/[the source path]` for
the subrepos.

Then there's a short summary that explains what the change does.

The summary should be the completion of the sentence "When committed, this
change will..."

Of all the parts of the change description this line is the most important part.


![](slides/anatomy_change_description_second_line.png)


After that, we provide more details about the change.

Finally, if the change has a related github issue (and it should unless the
change is really simple), we reference it.

![](slides/change_descriptions_gopherbot.png)

We've got some magic that links between changes and an issue. If a change
mentions an issue, gopherbot will post a comment on the change, and when
a fixing change is submitted, gopherbot will close the issue.

Just to stress the importance of the first line of the change description,
here's an example of what not to do.

![](slides/change_descriptions_what_not_to_do.png)

I was once hacking on a change and set
the change description to "i XXX" with the plan to go back and revise it before
sending it out. But I "git mail"ed it before fixing the subject line, and all
the emails sent out for the change had the subject "i XXX"..., even after I
changed the commit message. A bunch of people were annoyed by this.... And I had
to scrap the change and send a new one out that had a reasonable subject line...
Once you send out your change for review, the email subject line is fixed.

Ok. Back to the demo.

Let my write my description...

    regexp: add the Fanout benchmark

    This is a copy of the "FANOUT" benchmark recently added to RE2 with the
    following comment:

        // This has quite a high degree of fanout.
        // NFA execution will be particularly slow.

    Most of the benchmarks on the regexp package have very little fanout and
    are designed for comparing the regexp package's NFA with backtracking
    engines found in other regular expression libraries. This benchmark
    exercises the performance of the NFA on expressions with high fanout.

## Generating password

![](slides/generate_password.png)

Ok we're almost ready to get this change reviewed. Two more things:

First, we need to generate a cookie to communicate our identity to the codereview tool

I'm going to skip demoing this part since the page generates a key that'll let
you submit changes as me. But let's take a look at the process.

Go to go.googlesource.com and log in. Then click "generate password". That will take you to a
page with some code you can paste into your shell to add the key.

## CLA

![](slides/cla1.png)

![](slides/cla2.png)
We'd also need to file a
Contributor License Agreement granting the Go project a copyright license to
redistribute your code. git mail will ask you to do that if you haven't yet.
Here's how to get to the agreement... golang.org/cl Here's what it looks like.

Remember that your git login email should be the same as the email in your
`.gitconfig`.

Now, I've already got a CLA, but git mail still needs to know who I am before
letting me send out my change.

## Codereview

![](slides/code_review.png)

Ok.

The next step is to get our change code reviewed.

In the code review process, you and your reviewers (who are contributors like
you) work together to improve your change. 

The reviewers examine and approve every
change we make to our source code. If they think something needs to be changed,
they'll post a comment on your change, and you can respond to comments and make
changes in future revisions.

Multiple rounds of review are normal, even for small changes. And don't forget
to run gofmt and the tests every time you upload your review.

Reviewers will also use the gerrit scores we saw earlier to allow a change to be
submitted or to block it. These scores help ensure that there's agreement on
the change that's being made before it's submitted. 

Remember, they're scoring the change, it's not personal.

Let's mail our change out:

    git show

We use `git show` to look at the change before sending it out. If your change isn't
formatted, run `git gofmt` to format your change.

    git mail

Once your reviewer is happy with your change, they'll give you a '+2' on your
commit. That means they think your change is ready to be submitted.

And with all that done, your change will be ready to be submitted! While you're
a beginner contributor you'll have to ask your reviewer to submit your change
for you. (We'll give you committer rights once we're sick of submitting your
changes for you :) )


# What if no one responds?

When you're contributing, you might find that no one responds to your changes.
If that happens, your best bet is to send a friendly ping on the change.
The same goes to issue requests.

It happens that changes fall through the cracks -- we aren't deliberately ignoring
you! So just ping the change to surface it.

On the other hand because people are busy, you might need to wait a day or two
before getting a response to your change or comments. It's best to give your
reviewer _some_ time to respond before pinging.

# Advanced contributors

![](slides/advanced_contributors.png)

Okay, so we've talked about making simple contributions... let's talk about more
complex changes:

Any work more complicated than what we've already mentioned generally requires
an bug tracker issue that's been accepted, and in even more involved cases, a
proposal.

## The Go release cycle and the freeze

![](slides/release_cycle.png)

And another thing: we'll accept bug reports and proposals, doc changes and small tests and
benchmarks year round, but you should be mindful of the Go release cycle if
you're contributing code or doing anything more complex, especially on the core
go repository.

Each Go release is about 6 months long. The Go tree is open for development for
the first three months of the cycle, and enters the freeze for the second three
months. During the freeze, we'll generally only make changes to fix major bugs
or regressions introduced during the development phase. We're currently in the
freeze for the 1.7 release cycle. When 1.7 is released in August, the tree will
thaw and the 1.8 development phase will begin.

The tree will then be open for general contributions for three months, after
which the 1.8 freeze will begin.

If you're itching to make a change, you can still send it out and add the
comment "R=go1.8", but it won't be looked at and can't be submitted until the
thaw.

![](slides/r=go1.8.png)

Here's a R=go1.8 example. I didn't finish my regular expressions work in time, so we added
a comment with the text "r=go1.8" to indicate to our tools and contributors that
this change should be ignored till the thaw.

## Fixing accepted bugs

![](slides/fixing_bugs.png)

If you see an open issue with no assignee you should feel free to take it on.
Comment on the issue to let people know that you're working on it so we can
avoid having people do the same work twice. Then send in your change as before.

It's ok if you've volunteered to fix an issue, but find that you don't
have the time. But if that happens, remember to send another comment letting
people know that the change is open for a new assignee...

Once you have a fix, send it to us. Most changes should include tests,
especially if you're fixing a bug!

## Proposing features and changes

![](slides/proposals.png)

If you want to make a change introducing new behavior or features to the runtime
and compiler, standard libraries, tools, or subrepos, start by filing an
issue.

The project contributors will review the issue and make a decision to accept the
proposal, reject it, or ask for a proposal document. Discussion will continue on
the issue tracker until a decision is made.

Aside: The _language_ is more or less frozen and language changes must be small,
focused, and address a significant need in the ecosystem.

In contrast, much of the new work that's being done is more behind the scenes:
for example optimizations and bugfixes in the the standard library, GC
improvements in the runtime, refactoring and the SSA backend in the compiler,
and vendoring support in the tools and, there's a lot of work that's being done
in the subrepos.

## Tips for reviewers

![](slides/tips_for_reviewers.png)

And now, some tips for reviewers.

One of the most helpful ways to contribute to the Go project is by reviewing
changes. You can actually start reviewing changes even before you become a
committer. And even if you're not familiar with the codebase, there are a few
things most contributors can help point out in reviews:

*   Don't be afraid to say "I don't understand this function". It might cause
    the author to provide a good explanation in a comment, or better yet, to
    simplify the code. (We love simple code!)

*   You can suggest improvements to change descriptions. Remember that
    descriptions have a particular format

*   You can suggest code style improvements. All new code added to the go
    project must should follow the guidelines in 'Effective Go', and the style
    advice collected in the 'Go review comments doc' at golang.org/s/style.

*   All non-trivial changes should point back to the GitHub issue they're
    addressing. Proposals should be approved before work starts on them

It's okay to be picky; but when you are, acknowledge that you are. Sometimes
people feel like they're being blocked, when really you're just addressing a
small quibble that is ultimately not a big deal.

And don't forget, anyone can make review comments on changes. You don't need to
be a committer to help out with reviews.

# Wrapping up

![](slides/wrapping_up.png)

Our contributors come from many different backgrounds and have different
motivations for contributing to the project. So everyone's path is going to be
different... but let's take a step back and look at an example of a progression
of involvement on the project:

If there's a particular part of the project that you're interested in, start by
reading the code. If you're confused, that might be an opportunity! It might
mean that there are missing test cases or documentation. Look for bugs in the
github issue tracker that are relevant to that part and try to fix them. Once
you've fixed a number of issues and start becoming more familiar with the code
base, you might get some ideas for changes you want to make and things you want
to fix, so you might start making suggestions to the issue tracker or filing
proposals. And you might start watching for changes done in the directory and
adding review comments.

Eventually, we'll get tired of submitting your changes for you and we'll make
you a committer, so you can submit them yourself. But you'll also be able to
grant "+2"'s to other changes, approving them to be submitted, and to submit
changes for new contributors. At this point you might become active in reviewing
changes to that part of the codebase.

## Remember to be nice:

![](slides/remember_be_nice.png)

Remember: this is true for every contributor, but especially for advanced
contributors. You're representing the go community. Remember to be nice!

And of course,

## Thank you for contributing!

![](slides/thank_you.png)

*Thank you* to ALL our contributors for your contributions! Y'all make our
project great!

*Thank you* to ALL our reviewers for the time you spend reviewing. The work you
do in improving our code quality and teaching our contributors is invaluable!

And *Thank Y'all* for your interest in contributing. You're the future of Go!

# Questions?

![](slides/questions.png)