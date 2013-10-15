# rescat

A golang RESource conCATenator

## Purpose, History and Goals

The general purpose of go-rescat is to provide a go implementation of [perlbal's enable_concatenate_get](http://search.cpan.org/~dormando/Perlbal-1.80/lib/Perlbal/Manual/WebServer.pod#enable_concatenate_get) and [mod_concat](https://code.google.com/p/modconcat/). mod_concat is an Apache2 module to provide perlbal's enable_concatenate_get option.

### Purpose

The above modules improve website performance and web development efficiency by allowing independent developement of individual JS and CSS resources by different teams in different files, but then effectively serving them in a single request per resource type -- JS and CSS.

Given an example where mobile developers work in `http://example.com/css/mobile.css`, and desktop developers work in `http://example.com/css/default.css` and `http://example.com/css/admin.css` and all teams are building their work on a common CSS framework `http://example.com/css/framework/main.css`. One can effectively optimize the CSS loaded into a page into a single request -- optimized per page [note the per page optimization would still be manual unless automated by some build process].

For public facing pages, one might use the following resource URL `http://example.com/css/??framwork/main.css,default.css,mobile.css` rather than three individual stylesheet links. Further, for admin pages one might use ``http://example.com/css/??framwork/main.css,default.css,mobile.css,admin.css`. Similar optimizations might happen for JS, `http://example.com/js/??jquery.js,jqueryui.js,default.js`.

NB: The double `??` denoted to perlbal and mod_concat that they should implement concatenation behavior for this URL. The the order of the comma seperated resources is important in that it is the order in which the concatenation of resources. 

### History

TODO: Add some copy about the history of: perbal, mod_concat, Steve Souders [Rule 1](http://stevesouders.com/hpws/rule-min-http.php) from [High Performance Web Sites](http://stevesouders.com/hpws/rules.php) and why rescat exists...



### Goals

While this software may be moot with the advent and widespread adoption of SPDY or HTTP 2.0, it may still prove useful for some period of time until all agents support SPDY or HTTP 2.0. Additionally it serves as a well know problem with a (roughly) fixed scope, as well as sufficient complexity to leverage and learn a significant set of features in Go. 

This package is a library which implements an [http.Handle](http://golang.org/pkg/net/http/#Handle) or [http.HandleFunc](http://golang.org/pkg/net/http/#HandleFunc) with configurable providers usable within any generic web framework as well as provide examples of such integration.

Additionally, it provides an http server implementation which allows configuring multiple domains with multiple providers and handlers.

## Definitions

### Handlers

Handlers are [http.Handle](http://golang.org/pkg/net/http/#Handle) functions with can be used in a [ServeMux](http://golang.org/pkg/net/http/#ServeMux).

### Providers

TODO: elaborate

Ideas for providers include: 

  - filesystem provider: retrieve the requested resources on the filesystem at the configured document root/base path
  - http provider: retrieve the requested resources via http requests from the configured base URL 
  - S3 provider: retrieve the requested resources via http requests from the configured bucket
  - groupcache or memcache provider: ...

Ideas for server options/optimizations:
  - gzip compression
  - minification
  - cache/expiration header configuration
  - memoization
  - group cache/memcache/redis integration

<small>Copyright © 2013 Reed O'Brien. All rights reserved.</small>

