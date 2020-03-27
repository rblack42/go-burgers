go-burgers (v0.0.6)
###################
:Course: COC2325 - Computer Architecture and Machine Language
:School: Austin Community College
:Email: roie.black@gmail.com
:Documentation: http://rblack42.github.io/go-burgers/index.html

|travis-build| |license|

Computer Science students in most programs never have a chance to explore
parallel processing. Since I am both and Aerospace engineer, and a Computer
Scientist, I decided to build a simple tool that can be used to explore Burgers
Equation, commonly used in introducing Aerospace engineering students to the
world of Computational Fluid Dynamics. For CS students, the equation itself may
not be of much interest. However, looking at how engineers use computers to
solve important real-world problems will give them some insight into how we
write programs that really tax our computers to the max!

The code in this project is written in Go, which is a modern language well
suited for exploring parallel processing concepts. As the project evolves, it
will switch from running on a single processor to a cluster of processors,
simulating a modern high-performance computer system. Currently, the target
cluster is a 16-node Raspberry Pi 4 machine.

Dependencies
************

    * Go - tested using version 1.13

    * Fyne_  - GUI interface

    * Travis-CI_ - CI testing


..  _Travis-CI:     https://travis-ci.org/
..  _Fyne:  https://fyne.io/

..  |travis-build| image:: https://travis-ci.org/rblack42/go-burgers.svg?branch=master
    :alt: Build badge from Travis-CI

..  |license| image:: https://img.shields.io/badge/License-BSD%203--Clause-blue.svg
    :alt: BSD 3-Clause License









