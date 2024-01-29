import { Component } from '../component'
import './add.css'

class Add extends Component {
  async showAddDialog() {
    console.log('showAddDialog')
    const url = prompt('Enter a site or feed URL')
    if (!url) return
    try {
      const response = await fetch('/api/links', {
        method: 'post',
        body: JSON.stringify({ url }),
      })
      const json = await response.json()
      console.log(json)
    } catch (e) {
      console.error(e)
    }
  }
  connectedCallback() {
    this.addEventListener('click', () => this.showAddDialog())
  }
}
customElements.define('linky-add', Add)
