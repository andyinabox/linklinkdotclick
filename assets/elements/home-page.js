import { ProgressiveBody } from '../lib/progressive-body'
class HomePage extends ProgressiveBody {}
customElements.define('home-page', HomePage, { extends: 'body' })
