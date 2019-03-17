This is for an idea I had today. Open to finding some people to collaborate on it with me.

# What is FreeD?

FreeD is an idea I had for a new approach to 3d printing (at least I think it's new, I'm still in the early stages of vetting this idea out). It is based on a couple of existing sterolithography-based 3d printing technoligies that already exist. The idea is to use "slices" of a 3d mesh, and project them onto a viscous photo-curing resin. A projector would have a variable focus that would select the depth at which it would print the current layer, and then another plane laser would shine a plane of uv light along that same depth. The intersection of the projected image with the plane laser would have a high enough intensity to kick off the cross-linking reaction that cures the resin, and the viscous resin surrounding the print would eliminate the need for supports. I had an idea similar to this a while back, and recently stumbled across [this video](https://youtu.be/jcwYFBeetH0) from UC Berkeley that inspired me to try and improve on my design.

If this project doesn't interest you, or you feel like you have other ideas/problems you'd like to see being worked on, I recently purchased the domain crowdsrc.io, and it's under development right now. Ultimately I want it to be a place where we can involve people who may not have a programming background in the identification and solution of problems. Go check out the git repo (make a git repo and link it here) for a more in-depth description.

This module contains a few different go packages, the intent of each package is roughly outlined below.

# Parsers

As of right now, my plan is to initially support both STL and OBJ files, with STL taking a preference as that's the file format I see the majority of 3d printables being uploaded in. Basically all we need back from it is a collection of trinagles to pass to the slicer.

## STL Parser
*Package Name:* github.com/dkkempto/freed/parser/stl

This package is used to decode stl files into a format that is useful to us. I'm sure something like this already exists, but this project is also for me to start learning the ins and outs of go :)

## OBJ Parser
*Package Name:* github.com/dkkempto/freed/parser/obj

This package is used to decod .obj files.

# Slicer
*Package Name:* github.com/dkkempto/freed/slicer

This package is responsible for taking a collection of triangles, and providing a utlity to retrieve a collection of paths to be sent to the printing device (in our case this will most likely just be the renderer package which will just convert it into an image, but we could just as easly add support for another printing device like a more traditional 3d printer or cnc)

# Renderer
*Package Name:* github.com/dkkempto/freed/renderer

This package is responsible for taking the paths from our slicer and writing them to an image/sequence of images to be sent to the projector.