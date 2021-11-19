<script>
	import Stock from '../../artifacts/contracts/Stock.sol/Stock.json'
	import NumberOfOwners from './NumberOfOwners.svelte';
	import MyStocks from './MyStocks.svelte';
	import TotalStocks from './TotalStocks.svelte';
	import Owners from './Owners.svelte';
	import { ethers } from 'ethers'
	import { onMount } from 'svelte'
	import detectEthereumProvider from '@metamask/detect-provider'

	export let name, stockAddress;

</script>

<main>

	<div class="container">
		<div class="row justify-content-center">
			<h1>Decentralized {name}</h1>
		</div>
		<div class="row justify-content-center">
			<h4>Dapp for decentralized stock options</h4>
		</div>

		<div class="row justify-content-center mt-5">
	    <div class="col mt-2">
				<MyStocks stockAddress={stockAddress} />
	    </div>
	    <div class="col mt-2">
				<TotalStocks stockAddress={stockAddress} />
	    </div>
	    <div class="col mt-2">
				<NumberOfOwners stockAddress={stockAddress} />
	    </div>
	  </div>

		<div class="row justify-content-center mt-5">
			<h5>All owners and their stocks</h5>
		</div>

		<div class="row">
			<div class="col-12 mt-2">
	    	<Owners stockAddress={stockAddress} />
			</div>
	  </div>

	</div>

</main>

<style>
	main {
		text-align: center;
	}

	h1 {
		color: #5a57fa;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
