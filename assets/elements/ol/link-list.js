import { handleError } from '../../lib/errors'
import { createLink } from '../../lib/api'
import { eventsMixin } from '../../lib/mixins'

const FETCH_BUFFER = 10 * 60 * 1000 // 10 minutes
export class LinkList extends HTMLOListElement {
  constructor() {
    super()
    this.template = this.querySelector('template')
    this.fetchAllLinks()
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
    return Array.from(this.querySelectorAll('li'))
  }

  async fetchAllLinks() {
    const links = [...this.links]
    this.loading = true
    const promises = links.map((link) => link.fetchData())
    await Promise.all(promises)
    this.loading = false
    this.sortLinks()
    this.lastFetched = Date.now()
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
      this.sortLinks()
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  sortLinks() {
    this.loading = true
    const links = [...this.links]
    links.sort((a, b) => {
      const d1 = new Date(a.data.lastClicked).getTime()
      const d2 = new Date(b.data.lastClicked).getTime()
      if (d1 === d2) {
        return 0
      }
      return d1 < d2 ? -1 : 1
    })
    this.innerHTML = ''
    links.forEach((link) => this.appendChild(link))
    this.loading = false
  }

  connectedCallback() {
    this.listen(document, 'link-create-request', this.createLink)
    this.listen(this, 'link-click-success', this.sortLinks)
    this.listen(this, 'link-update-success', this.sortLinks)
    this.listen(window, 'focus', () => {
      if (!this.lastFetched || Date.now() - this.lastFetched > FETCH_BUFFER) {
        this.fetchAllLinks()
      }
    })
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}

Object.assign(LinkList.prototype, eventsMixin)
customElements.define('link-list', LinkList, { extends: 'ol' })
