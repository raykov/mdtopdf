# H1
## H2
### H3
#### H4
##### H5
###### H6

Alternatively, for H1 and H2, an underline-ish style:

Alt-H1
======

Alt-H2
------

---

Emphasis, aka italics, with *asterisks* or _underscores_.

Strong emphasis, aka bold, with **asterisks** or __underscores__.

Combined emphasis with **asterisks and _underscores_**.

Strikethrough uses two tildes. ~~Scratch this.~~

---

1. First ordered list item
2. Another item
  * Unordered sub-list.
1. Actual numbers don't matter, just that it's a number
  1. Ordered sub-list
4. And another item.

   You can have properly indented paragraphs within list items. Notice the blank line above, and the leading spaces (at least one, but we'll use three here to also align the raw Markdown).

   To have a line break without a paragraph, you will need to use two trailing spaces.  
   Note that this line is separate, but within the same paragraph.  
   (This is contrary to the typical GFM line break behaviour, where trailing spaces are not required.)

* Unordered list can use asterisks
- Or minuses
+ Or pluses

---

[I'm an inline-style link](https://www.google.com)

[I'm an inline-style link with title](https://www.google.com "Google's Homepage")

[I'm a reference-style link][Arbitrary case-insensitive reference text]

[I'm a relative reference to a repository file](../blob/master/LICENSE)

[You can use numbers for reference-style link definitions][1]

[I'm a link with ampersands](https://www.google.com?some=information&another[]=list&another[]=information)

Or leave it empty and use the [link text itself].

URLs and URLs in angle brackets will automatically get turned into links.
http://www.example.com or <http://www.example.com> and sometimes
example.com (but not on Github, for example).

Some text to show that the reference links can follow later.

[arbitrary case-insensitive reference text]: https://www.mozilla.org
[1]: http://slashdot.org
[link text itself]: http://www.reddit.com


Here's our logo (hover to see the title text):

Inline-style:
![alt text](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Logo Title Text 1")

Reference-style:
![alt text][logo]

[logo]: https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Logo Title Text 2"

Base64-style:
![alt text](data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAYAAABXAvmHAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAABmJLR0QAAAAAAAD5Q7t/AAAACXBIWXMAAG66AABuugHW3rEXAAAAB3RJTUUH4QIKFjcfARWjgAAAAuVJREFUaN7tmT1IHEEYhtdfBOMpiEjujgtWNoKdrU0EMfGiRYKFVkJiOBU0FiEmhZWVpWACJ1EQWxEEISb+XaGFxEiMqJWFVgmIcJ16eQZm4Thcd2939ifhBp7GO9f3kZ35Zr7RtHvG6/DHFDyDIs2H8S42J4hDytYDCJ6Ba1iFdq9ECFwE7bAK15BxIpDNGnS4JSKDd8CaCJ2NKgGdb/AUihUFL4Yn8DU3uFsCOuvQaVdEBu+EdaPgbgvobMjJXmIxeImcnBtmwb0S0NmCLiMRGbwLtqwG91pAZxu6E+FPpeI5b2OfSwnRDdv5BvdLQCc1Ek2+cRLcV4EBQCAz7jB8QeB/FrhVJHDrl4DYWgzCqU2BUxi8a+vglcCS/F4NDMGxRYFjGIIaWR+W/BJYzvl+NSTgyEDgCBJQnVPglgMhkPV7IRhA4FAK/CLEAIQMKnSwBPTx8uFM1Wg0+fjDo/kqk71RMAXy2JUWBAoCBYGCQEEgmAJ9sOOzwA702X44IcuhB3Y9FtiFHihX0jUzEHFDQF1wAjbcI7IHK6CigbsCe0bB+VmDXYE/MAn1d3xWAc1O+6SyH9oMFXd8Vg+T8NvpKnQmDyyVHnWnK+WB50z1MroPL1Q1dQ16pc9h3+06IM7CrYrDt8IX1XXg5z0S4uJjAZocBm+CBf0iw4BDuwK1MGtSja9gCiJ5Bo/AFFyZVOFZqHW6nA5D2kTkHMbEedgkeAjG4NwkeBqGlb2j4p2HEwv7o0PohdKc4KJD3SteBwuduBMxJ5SvEuI1EdXXYmtxE9pGoskywrTBpsU2oqjKEdfWaPGfhQk5ic16ozcIfB+Pzd1YCC4m8QSUeXKHK29iLhT1Ri/EzY3nF9GEbJS3MU4ExOVHo+bXENsLmLYpMA0PtCAMAvfDpUWBS+jXgjYI3QIHJgIH0KIFdRC8DhYNBBahTgv6EGcEUZURSEuBtKy+Rdq/NF6FZ+Kj0eSP97H5uFt/4y9GA721POPJpAAAAABJRU5ErkJggg== "Logo Title Text 1")

---

Inline `code` has `back-ticks around` it.

```javascript
var s = "JavaScript syntax highlighting";
alert(s);
```

```python
s = "Python syntax highlighting"
print s
```

```
No language indicated, so no syntax highlighting. 
But let's throw in a <b>tag</b>.
```


---

Colons can be used to align columns.

| Tables        | Are           | Cool  |
| ------------- |:-------------:| -----:|
| col 3 is      | right-aligned | $1600 |
| col 2 is      | centered      |   $12 |
| zebra stripes | are neat      |    $1 |

There must be at least 3 dashes separating each header cell.
The outer pipes (|) are optional, and you don't need to make the
raw Markdown line up prettily. You can also use inline Markdown.

Markdown | Less | Pretty
--- | --- | ---
*Still* | `renders` | **nicely**
1 | 2 | 3

---

> Blockquotes are very handy in email to emulate reply text.
> This line is part of the same quote.

Quote break.

> This is a very long line that will still be quoted properly when it wraps. Oh boy let's keep writing to make sure this is long enough to actually wrap for everyone. Oh, you can *put* **Markdown** into a blockquote.

---

<dl>
  <dt>Definition list</dt>
  <dd>Is something people use sometimes.</dd>

  <dt>Markdown in HTML</dt>
  <dd>Does *not* work **very** well. Use HTML <em>tags</em>.</dd>
</dl>


---

Three or more...

---

Hyphens

***

Asterisks

___

Underscores

