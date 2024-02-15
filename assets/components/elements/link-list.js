import { handleError } from '../../lib/errors'
import { createLink } from '../../lib/api'
import { eventsMixin } from '../../lib/mixins'
export class LinkList extends HTMLOListElement {
  constructor() {
    super()
    this.template = this.querySelector('template')
  }

  get loading() {
    return this.classList.hasClass('loading')
  }
  set loading(v) {
    if (v) {
      this.classList.add('loading')
      this.broadcast('loading-start')
    } else {
      this.classList.remove('loading')
      this.broadcast('loading-stop')
    }
  }

  get links() {
    return Array.from(this.querySelector('li'))
  }
  async createLink() {
    let url = prompt('Enter a website or feed URL')
    if (!url) return

    try {
      this.loading = true
      const link = await createLink(url)
      const linkEl = this.template.content.firstElementChild.cloneNode(true)
      this.prepend(linkEl)
      linkEl.data = link
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  sortLinks() {
    this.broadcast('loading-start')
    const links = [...this.links]
    links.sort((a, b) => {
      const d1 = new Date(a.data.lastClicked).getTime()
      const d2 = new Date(b.data.lastClicked).getTime()
      if (d1 === d2) {
        return 0
      }
      return d1 < d2 ? -1 : 1
    })
    const linksContainer = this.slots.links
    linksContainer.innerHTML = ''
    links.forEach((link) => linksContainer.appendChild(link))
    this.broadcast('loading-stop')
  }

  connectedCallback() {
    this.listen(document, 'link-create-request', this.createLink)
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}

Object.assign(LinkList.prototype, eventsMixin)
customElements.define('link-list', LinkList, { extends: 'ol' })
