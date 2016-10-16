# htmlfixer

This is a console app written in go (golang). 

## Background
This application is used to help create html for inclusion in digital publications, specifically *epub* files.
 
One technique for creating digital publications is to write the text using MS-Word.  Word has a feature which allows one to save the document and convert it to html. (File -> SaveAs -> WebPage, Filtered).

The html created by MS-Word is problematic from the standpoint of creating digital publications.  This program, *htmlfixer*, will read single or multiple html files and convert them in place to more compatible css and html formats. 

Thereafter, the converted html files may be imported into a digital e-book editor like that provided by [Calibre](http://calibre-ebook.com/) for final publication.

## Usage 
*htmlfixer* extracts file path names for target html files using a text file named **FileList.txt** which must be located in the same directory as the *htmlfixer* executable.  Include full volume, path and file names in this input file.  Each file name must be placed on a separate line. Lines must be delimited by the new line character '\n' 

## History
Version 1.0 - This console app has been tested on Windows. Be advised the original html file(s) specified in source file **FileList.txt** will be deleted and replaced with the converted html files.