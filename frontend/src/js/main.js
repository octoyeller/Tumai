import '../css/style.css';
import '../css/app.css';

import { Exit_app } from '../../wailsjs/go/main/App';

document.addEventListener('keydown', (e) => {
  if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === 'q') {
    e.preventDefault();
    Exit_app();
  }
})
