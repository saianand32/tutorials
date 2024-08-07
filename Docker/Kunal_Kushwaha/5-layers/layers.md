
When u pull an image u will see it downloading as layers with hash values.. this is because multiple images might be already be using these layers so we can skip downloading the layers to speed up...

to check number of layers in a n image  -- 

$ docker images
copy the ID
 
$ docker inspect <ID>