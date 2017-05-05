## 附录A：[原文勘误](http://www.gopl.io/errata.html)

p.xiii, ¶1: delete "Object Oberon". (Thanks to Luca Boasso, 2016-02-15. Corrected in the fourth printing.)

p.9, ¶2: For "can compared", read "can be compared". (Thanks to Antonio Macías Ojeda, 2015-10-22. Corrected in the second printing.)

p.13: As printed, the gopl.io/ch1/lissajous program is deterministic, not random. We've added the statement below to the downloadable program so that it prints a pseudo-random image each time it is run. (Thanks to Randall McPherson, 2015-10-19.)

rand.Seed(time.Now().UTC().UnixNano())
p.15, ¶2: For "inner loop", read "outer loop". (Thanks to Ralph Corderoy, 2015-11-28. Corrected in the third printing.)

p.19, ¶2: For "Go's libraries makes", read "Go's library makes". (Thanks to Victor Farazdagi, 2015-11-30. Corrected in the third printing.)

p.31, ¶6: for "then the short variable declaration acts like an assignment to those variables", read "then, for those variables, the short variable declaration acts like an assignment" and in ¶7, the word "assigns" should be italicized for emphasis. (Thanks to Terry McKenna, 2016-11-18.)

p.40, ¶4: For "value of the underlying type", read "value of an unnamed type with the same underlying type". (Thanks to Carlos Romero Brox, 2015-12-19. Corrected in the fourth printing.)

p.40, ¶1: The paragraph should end with a period, not a comma. (Thanks to Victor Farazdagi, 2015-11-30. Corrected in the third printing.)

p.43, ¶3: Import declarations are explained in §10.4, not §10.3. (Thanks to Peter Jurgensen, 2015-11-21. Corrected in the third printing.)

p.48: f.ReadByte() serves as an example of a reference to f, but *os.File has no such method. For "ReadByte", read "Stat", four times. (Thanks to Peter Olsen, 2016-01-06. Corrected in the fourth printing.)

p.52, ¶2: For "an synonym", read "a synonym", twice. (Corrected in the second printing.)

p.52, ¶9: For "The integer arithmetic operators", read "The arithmetic operators". (Thanks to Yoshiki Shibata, 2015-12-20. Corrected in the fourth printing.)

p.57, ¶4: To "any comparison with NaN always yields false", add "(except !=, which is always the negation of ==)". (Thanks to Luca Boasso, 2016-02-15. Corrected in the fourth printing.)

p.68: The table of UTF-8 encodings is missing a bit from each first byte. The corrected table is shown below. (Thanks to Akshay Kumar, 2015-11-02. Corrected in the second printing.)

0xxxxxxx                             runes 0−127     (ASCII)
110xxxxx 10xxxxxx                    128−2047        (values <128 unused)
1110xxxx 10xxxxxx 10xxxxxx           2048−65535      (values <2048 unused)
11110xxx 10xxxxxx 10xxxxxx 10xxxxxx  65536−0x10ffff  (other values unused)

p.73, ¶1: For "a exercise", read "an exercise". (Thanks to vrajmohan, 2015-12-28. Corrected in the fourth printing.)

p.74: The comment in gopl.io/ch3/printints should say fmt.Sprint, not fmt.Sprintf. (Corrected in the second printing.)

p.75, ¶4: For "%u", read "%o". (Thanks to William Hannish, 2015-12-21. Corrected in the fourth printing.)

p.76: The comment // "time.Duration 5m0s should have a closing double-quotation mark. (Corrected in the second printing.)

p.79, ¶4: "When an untyped constant is assigned to a variable, as in the first statement below, or appears on the right-hand side of a variable declaration with an explicit type, as in the other three statements, ..." has it backwards: the first statement is a declaration; the other three are assignments. (Thanks to Yoshiki Shibata, 2015-11-09. Corrected in the third printing.)

p.84, ¶2: For "used as function parameters", read "used as function parameters or results". (Thanks to Yang Yang 杨玚, 2016-04-18. Corrected in the fourth printing.)

p.107, ¶6: in "JSON's \Uhhhh numeric escapes", the U should be lowercase: \uhhhh. (Thanks to Kohei Marumoto (丸本耕平), 2016-09-23. Corrected in the fourth printing.)

p.112: Exercise 4.11 calls for a "CRUD" (create, read, update, delete) tool for GitHub Issues. Since GitHub does not currently allow Issues to be deleted, for "delete", read "close". (Thanks to Yoshiki Shibata, 2016-01-18. Corrected in the fourth printing.)

p.115: The anchor element in gopl.io/ch4/issueshtml's template is missing a closing </a> tag. (Thanks to Taj Khattra, 2016-01-19. Corrected in the fourth printing.)

p.132, code display following ¶3: The final comment should read:
// compile error: can't assign func(int, int) int to func(int) int
(Thanks to Toni Suter, 2015-11-21. Corrected in the third printing.)

p.139, ¶4: For "in Chapter 3", read "in Chapter 4". (Thanks to Yoshiki Shibata, 2016-07-03. Corrected in the fourth printing.)

p.160, ¶4: For Get("item")), read Get("item"). (Thanks to Yoshiki Shibata, 2016-02-01. Corrected in the fourth printing.)

p.166, ¶2: For "way", read "a way". (Corrected in the third printing.)

p.177, ¶6: For "as we saw", read "as we will see". (Thanks to David Drysdale, 2016-07-05. Corrected in the fourth printing.)

p.200, TestEval function: The format string in the final call to t.Errorf should format test.env with %v, not %s. (Thanks to Mitsuteru Sawa, 2015-12-07. Corrected in the third printing.)

p.222, Exercise 8.1: The port numbers for London and Tokyo should be swapped in the final command to match the earlier commands. (Thanks to Kiyoshi Kamishima, 2016-01-08. Corrected in the fourth printing.)

p.235, ¶3: makeThumbnails refers to the second draft of the function, makeThumbnails2. (Thanks to 有賀功 Isao Aruga, 2016-07-13. Corrected in the fourth printing.)

p.272, ¶3: for "the request body", read "the response body". (Thanks to 曹春晖, 2016-01-19. Corrected in the fourth printing.)

p.288, code display following ¶4: In the import declaration, for "database/mysql", read "database/sql". (Thanks to Jose Colon Rodriguez, 2016-01-09. Corrected in the fourth printing.)

p.320, code display following ¶1: The suggested test case does not exercise the unary plus operator, nor is the eval package qualifier appropriate:

{"-x * -x", eval.Env{"x": 2}, "4"}
It should read:
{"+x * -x", Env{"x": 2}, "-4"}

(Thanks to Shinji Dosaka, 2016-07-15. Corrected in the fourth printing.)

p.341, Exercise 12.7: For "decoder" and "json.Decoder", read "encoder" and "json.Encoder", respectively. (Thanks to Yoshiki Shibata, 2016-08-27. Corrected in the fourth printing.)

p.347, Exercise 12.8: For "like json.Marshal", read "like json.Unmarshal". (Thanks to chai2010, 2016-01-01. Corrected in the fourth printing.)

p.362: The gopl.io/ch13/bzip program does not comply with the proposed rules for passing pointers between Go and C code because the C function bz2compress temporarily stores a Go pointer (in) into the Go heap (the bz_stream variable). The bz_stream variable should be allocated, and explicitly freed after the call to BZ2_bzCompressEnd, by C functions. (Thanks to Joe Tsai, 2015-11-18. Corrected in the third printing.)

p.366, Exercise 13.3: For bzip2.writer, read bzip.writer. (Thanks to Yoshiki Shibata, 2016-09-05. Corrected in the fourth printing.)
