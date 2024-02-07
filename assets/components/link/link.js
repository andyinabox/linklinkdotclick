import { getLink, updateLink, deleteLink } from '../../lib/api'
import { handleError } from '../../lib/errors'
import { Component } from '../component'
export class Link extends Component {
  constructor() {
    super()
    // this.fetchData()
  }

  // render data to the element
  render() {
    const data = this.data
    this.slots.link.href = data.siteUrl
    this.slots.link.textContent = data.siteName
    this.slots.count.textContent = data.unreadCount ? data.unreadCount : ''
    this.setAttribute('data-id', data.id)
  }

  async fetchData() {
    const id = this.getAttribute('data-id')
    if (!id) return
    try {
      this.loading = true
      this.data = await getLink(id)
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  async handleClick(e) {
    const link = this.data
    link.lastClicked = new Date().toJSON()
    try {
      const updatedLink = await updateLink(link)
      this.data = updatedLink
      this.dispatchEvent(new CustomEvent('link-click', { bubbles: true }))
    } catch (err) {
      handleError(err)
    }
  }

  async handleDelete(e) {
    const { id, siteName } = this.data

    if (!confirm(`Delete link "${siteName}"?`)) return

    try {
      this.loading = true
      await deleteLink(id)
      this.remove()
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  async handleRename(e) {
    const data = this.data
    const name = prompt(`Enter a new name for "${data.siteName}"`)
    if (!name) return
    data.siteName = name
    try {
      this.loading = true
      const updated = await updateLink(data)
      this.data = updated
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  connectedCallback() {
    this.onClick = (e) => this.handleClick(e)
    this.slots.link.addEventListener('click', this.onClick)

    const edit = this.slots['edit-menu']
    this.onDelete = (e) => this.handleDelete(e)
    edit
      .querySelector('[name="delete"]')
      .addEventListener('click', this.onDelete)
    this.onRename = (e) => this.handleRename(e)
    edit
      .querySelector('[name="rename"]')
      .addEventListener('click', this.onRename)
  }

  disconnectedCallback() {
    this.slots.link.removeEventListener('click', this.onClick)
    this.slots['edit-menu']
      .querySelector('[name="delete"]')
      .removeEventListener('click', this.onDelete)
  }
}

Link.create = function (parentEl, data) {
  const linkTmpl = document.getElementById('tmpl-link')
  const linkEl = linkTmpl.content.firstElementChild.cloneNode(true)
  parentEl.prepend(linkEl)
  linkEl.data = data
  return linkEl
}

customElements.define('linky-link', Link)
