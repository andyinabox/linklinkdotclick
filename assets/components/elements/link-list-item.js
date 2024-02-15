import { eventsMixin, slotsMixin } from '../../lib/mixins'
import { updateLink, deleteLink, clickLink, getLink } from '../../lib/api'
import { handleError } from '../../lib/errors'
class Link extends HTMLLIElement {
  constructor() {
    super()
    this.registerSlots()

    // replace link href
    const href = this.slots.link.getAttribute('data-href')
    if (href) {
      this.siteUrl = href
      this.removeAttribute('data-href')
    }

    if (this.linkId) {
      this.fetchData()
    }
  }
  async fetchData() {
    try {
      this.loading = true
      const link = await getLink(this.linkId)
      this.data = link
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }
  get data() {
    return { ...this._data }
  }
  set data(d) {
    this._data = d
    this.render()
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

  set linkId(v) {
    this.setAttribute('data-id', v)
  }
  get linkId() {
    return parseInt(this.getAttribute('data-id'))
  }
  set siteUrl(v) {
    this.slots.link.href = v
  }
  get siteUrl() {
    return this.slots.link.href
  }
  set siteName(v) {
    this.slots.link.innerText = v
  }
  get siteName() {
    return this.slots.link.innerText
  }
  set unreadCount(v) {
    this.slots.count.innerText = v
  }
  get unreadCount() {
    return parseInt(this.slots.count.innerText)
  }
  render() {
    const data = this.data
    if (!data) return
    const { id, siteName, unreadCount, siteUrl, hideUnreadCount, feedUrl } =
      data

    console.log('reander link', this.data)

    if (hideUnreadCount) {
      this.slots.count.textContent = ''
    } else if (unreadCount) {
      this.slots.count.textContent = data.unreadCount
    } else if (!feedUrl) {
      this.slots.count.textContent = '?'
    } else {
      this.slots.count.textContent = ''
    }

    this.linkId = id
    this.siteName = siteName
    this.siteUrl = siteUrl
    this.slots.form.data = data
  }
  async onClick() {
    try {
      this.loading = true
      const link = await clickLink(this.linkId)
      this.data = link
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }
  async onUpdate() {
    try {
      this.loading = true
      const link = await updateLink(
        Object.assign(this.data, this.slots.form.formData)
      )
      this.data = link
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }
  async onDelete() {
    if (!confirm(`Are you sure you want to delete ${this.data.siteName}?`))
      return
    try {
      this.loading = true
      const result = await deleteLink(this.data.id)
      this.remove()
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  connectedCallback() {
    this.listen(this.slots.link, 'click', this.onClick)
    this.listen(this.slots.form, 'link-update-request', this.onUpdate)
    this.listen(this.slots.form, 'link-delete-request', this.onDelete)
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
Object.assign(Link.prototype, eventsMixin, slotsMixin)
customElements.define('link-list-item', Link, { extends: 'li' })
