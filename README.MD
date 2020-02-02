## Download
[Latest release](https://github.com/Sierra410/sm-tool/releases)

# Part Manager

**BETA! Back up your mods before using it!**

Features:
- Description and title editor (supports all languages)
- Parts can be rearranged
- Can generate/edit part UUID + validate it
- Simple json (hjson) editor with error info

Mod directory can be passed as an argument:  
```smmodtool.exe -load=/path/to/my/mod```

![Screenshot](https://raw.githubusercontent.com/Sierra410/sm-tool/master/screenshot.png)

# Texture Compiler

Usage: ./smtcmp [DIR]...

If no DIR(s) is(are) provided current working directory is used instead.

The program will try to find all files that follow a specific naming pattern and combine them into two files: xxxx_dif.png and xxxx_asg.png

ASG group:

- xxxx_a.png   (Alpha)
- xxxx_s.png   (Specular)
- xxxx_g.png   (Glow)
- xxxx_r.png   (Reflectivity)

DIF group: 
- xxxx_d.png   (Diffuse)
- xxxx_da.png  (Diffuse alpha)
- xxxx_ao.png  (Ambient Occlusion)

Note:  
  All files must be in PNG format.  
  If there there's at least one file in any of the groups both DIF and ASG textures will be generated.  
  If there are no textures in a group an empty image with default resolution (1024x1024px) will be generated instead.  
  All files in the same group must have the same dimensions and the dimensions must be a power of 2!  
