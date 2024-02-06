import { getSelf } from '../../lib/api'
import { handleError } from '../../lib/errors'
import { Component } from '../component'
import { Link } from '../link/link'

export class Site extends Component {
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
  sortLinks() {
    const linksContainer = this.slots.links
    const links = linksContainer.querySelectorAll('linky-link')
    for (var i = 1; i <= links.length; i++) {
      const l1 = links[i - 1]
      const l2 = links[i]
      const d1 = new Date(l1.data.lastClicked)
      const d2 = new Date(l2.data.lastClicked)

      // swap links
      if (d2 > d1) {
        linksContainer.replaceChild(l2, l1)
        linksContainer.insertBefore(l1, l2)
      }
    }
  }
  async handleCreateLink() {
    try {
      const url = prompt('Enter a website or feed URL')

      if (!url) return

      this.loading = true
      const link = await createLink(url)
      Link.create(linksContainerEl, link)
      this.sortLinks()
    } catch (err) {
      handleError(err)
    } finally {
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

  render() {
    const data = this.data
    this.slots['site-title'].innerText = data.siteTitle
    document.head.querySelector('title').innerText = data.siteTitle
  }
  connectedCallback() {
    this.onRenameSiteClick = () => this.handleRenameSiteClick()
    this.slots['rename-site'].addEventListener('click', this.onRenameSiteClick)
    this.onEditClick = () => this.handleEditButtonClick()
    this.slots['edit'].addEventListener('click', this.onEditClick)
    this.onAddClick = () => this.handleCreateLink()
    this.slots['add'].addEventListener('click', this.onAddClick)
  }
  disconnectedCallback() {
    this.slots['rename-site'].removeEventListener(
      'click',
      this.onRenameSiteClick
    )
    this.slots['edit'].removeEventListener('click', this.onEditClick)
    this.slots['add'].removeEventListener('click', this.onAddClick)
  }
}
customElements.define('linky-site', Site)
