import { getSelf, updateSelf, createLink } from '../../lib/api'
import { handleError } from '../../lib/errors'
import { Component } from '../component'
import { Link } from '../link/link'

export class HomePage extends Component {
  constructor() {
    super()
    this.fetchData()
  }
  set editing(bool) {
    if (bool) {
      this.classList.add('editing')
    } else {
      this.classList.remove('editing')
    }
  }
  get editing() {
    return this.classList.contains('editing')
  }
  get links() {
    return Array.from(this.slots.links.querySelectorAll('linky-link'))
  }
  async fetchData() {
    try {
      this.loading = true
      const self = await getSelf()
      this.data = self
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  render() {
    const data = this.data
    this.slots['site-title'].innerText = data.siteTitle
    document.head.querySelector('title').innerText = data.siteTitle
  }

  async reloadAllLinks() {
    try {
      this.loading = true
      await Promise.all(this.links.map((link) => link.fetchData()))
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
      const d1 = new Date(a.data.lastClicked)
      const d2 = new Date(b.data.lastClicked)
      return d1 < d2 ? -1 : 1
    })
    const linksContainer = this.slots.links
    linksContainer.innerHTML = ''
    links.forEach((link) => linksContainer.appendChild(link))
    this.loading = false
  }

  async handleCreateLink() {
    try {
      let url = prompt('Enter a website or feed URL')

      if (!url) return

      if (url.indexOf('http') !== 0) {
        url = 'http://' + url
      }

      this.loading = true
      const link = await createLink(url)
      Link.create(this.slots.links, link)
    } catch (err) {
      handleError(err)
    } finally {
      this.sortLinks()
      this.loading = false
    }
  }

  async handleRenameSiteClick() {
    try {
      const siteTitle = prompt('Enter a new title')

      if (!siteTitle) return

      this.loading = true
      const user = await updateSelf({ siteTitle })
      this.data = user
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  handleEditButtonClick() {
    const button = this.slots.edit
    if (this.editing) {
      button.textContent = 'Edit'
      this.editing = false
    } else {
      button.textContent = 'Done'
      this.editing = true
    }
  }

  connectedCallback() {
    this.reloadAllLinks() // re-fetch all links
    this.listen(this.slots['rename-site'], 'click', this.handleRenameSiteClick)
    this.listen(this.slots['edit'], 'click', this.handleEditButtonClick)
    this.listen(this.slots['add'], 'click', this.handleCreateLink)
    this.listen(this, 'link-click', this.sortLinks)
    this.listen(window, 'focus', this.reloadAllLinks)
  }
}
customElements.define('linky-home-page', HomePage)
