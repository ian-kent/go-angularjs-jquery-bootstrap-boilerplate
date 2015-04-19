Go, AngularJS, jQuery and Bootstrap
===================================

Boilerplate Go web application with:

* [AngularJS 1.3.15](https://angularjs.org/)
* [jQuery 2.1.3](https://jquery.com/download/)
* [Bootstrap 3.3.4](http://getbootstrap.com/)
* [gorilla/pat](https://github.com/gorilla/pat)
* [unrolled/render](https://github.com/unrolled/render)
  - Currently uses forked [ian-kent/render](https://github.com/ian-kent/render) to
    support loading [go-bindata](https://github.com/jteeuwen/go-bindata) compiled assets.

### Template delimiters

Go templates use `{{` and `}}` by default, which is incompatible with AngularJS syntax.

This boilerplate replaces the default delimiters with `[:` and `:]`.

### Assets

Assets are compiled using [go-bindata](https://github.com/jteeuwen/go-bindata).

- `assets/static` files are served as static files.
- `assets/templates` with suffix `.tmpl` are compiled and rendered by [unrolled/render](https://github.com/unrolled/render)

## License

Copyright ©‎ 2015, Ian Kent (http://iankent.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
