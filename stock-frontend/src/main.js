import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
		name: 'stock',
		stockAddress: '0x5FbDB2315678afecb367f032d93F642f64180aa3'
	}
});

export default app;
