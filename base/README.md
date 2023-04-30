# Base

Create default collections:

- bookmark_services
    - name: string
- bookmark_items
    - name: string
    - url: url
    - service: bookmark_services(single)

Create some services as only needed

- GitHub
- Confluence
- etc

## Import data

Import existing bookmarks, open the web console and,

```js
const data="<the json data>"
data.items.forEach(b => {
    fetch("/api/collections/bookmark_items/records", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            "name": b.title,
            "url": b.arg,
            "service": "52uvopuvkz7yt3m"
        })
    })
})
```
