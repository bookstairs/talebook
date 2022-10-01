# Talebook (Golang)

The Golang implementation for [talebook](https://github.com/talebook/talebook).
We only reuse the frontend code and rewrite backend in Golang.

In order to make it compatible with the original talebook, we will keep using the calibre as the book manager.
So you have to deploy this talebook fork in docker. But this will make it easy to fallback use the Calibre application.

## Talebook API implementation status

### Books API

- [ ] `/api/index`
- [ ] `/api/search`
- [ ] `/api/recent`
- [ ] `/api/hot`
- [ ] `/api/book/nav`
- [ ] `/api/book/upload`
- [ ] `/api/book/([0-9]+)`
- [ ] `/api/book/([0-9]+)/delete`
- [ ] `/api/book/([0-9]+)/edit`
- [ ] `/api/book/([0-9]+)\.(.+)`
- [ ] `/api/book/([0-9]+)/push`
- [ ] `/api/book/([0-9]+)/refer`
- [ ] `/read/([0-9]+)`

### Users API

- [ ] `/api/welcome`
- [ ] `/api/user/info`
- [ ] `/api/user/messages`
- [ ] `/api/user/sign_in`
- [ ] `/api/user/sign_up`
- [ ] `/api/user/sign_out`
- [ ] `/api/user/update`
- [ ] `/api/user/reset`
- [ ] `/api/user/active/send`
- [ ] `/api/active/(.*)/(.*)`
- [ ] `/api/done/`

### Metadata API

- [ ] `/api/(author|publisher|tag|rating|series)`
- [ ] `/api/(author|publisher|tag|rating|series)/(.*)`
- [ ] `/api/author/(.*)/update`
- [ ] `/api/publisher/(.*)/update`

### Files API

- [ ] `/get/pcover`
- [ ] `/get/progress/([0-9]+)`
- [ ] `/get/extract/(.*)`
- [ ] `/get/(.*)/(.*)`
- [ ] `/(.*)`

### ADMIN API

- [ ] `/api/admin/ssl`
- [ ] `/api/admin/users`
- [ ] `/api/admin/install`
- [ ] `/api/admin/settings`
- [ ] `/api/admin/testmail`
- [ ] `/api/admin/book/list`

### Scan API

- [ ] `/api/admin/scan/list`
- [ ] `/api/admin/scan/run`
- [ ] `/api/admin/scan/status`
- [ ] `/api/admin/scan/delete`
- [ ] `/api/admin/scan/mark`
- [ ] `/api/admin/import/run`
- [ ] `/api/admin/import/status`

### OPDS API

- [ ] `/opds/?`
- [ ] `/opds/nav/(.*)`
- [ ] `/opds/category/(.*)/(.*)`
- [ ] `/opds/categorygroup/(.*)/(.*)`
- [ ] `/opds/search/(.*)`

## Usage

TODO
