import { createLink } from '../lib/api'
import { eventsMixin } from '../lib/mixins'
class HomePage extends HTMLBodyElement {
  constructor() {
    super()
  }
  async handleLinkAddRequest(evt) {
    const url = evt.detail.url
    try {
      this.loading = true
      const link = await createLink(url)
      this.broadcast('link-add-success', { link })
    } catch (err) {
      handleError(err)
    } finally {
      this.sortLinks()
      this.loading = false
    }
  }
  connectedCallback() {
    this.listen(this, 'link-add-request', (evt) => console.log('link-add'))
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
Object.assign(HomePage.prototype, eventsMixin)
customElements.define('linky-home-page', HomePage, { extends: 'body' })
