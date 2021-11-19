<script>
	import Stock from '../../artifacts/contracts/Stock.sol/Stock.json'
	import {
		Col,
		Container,
		Row,
		Button } from 'sveltestrap';
  import { getContext } from 'svelte';
	import NumberOfOwners from './NumberOfOwners.svelte';
	import MyStocks from './MyStocks.svelte';
	import TotalStocks from './TotalStocks.svelte';
	import Owners from './Owners.svelte';
	import { ethers } from 'ethers'
	import { onMount } from 'svelte'
	import detectEthereumProvider from '@metamask/detect-provider'

	export let name, stockAddress, addressReceiver = '', amountStocksToSend = 0;

  async function requestAccount() {
    await window.ethereum.request({ method: 'eth_requestAccounts' });
  }

  async function transferStocks() {
    if (typeof window.ethereum !== 'undefined') {
      await requestAccount()
      const provider = new ethers.providers.Web3Provider(window.ethereum)
      const signer = provider.getSigner()
      console.log({ provider })
      const contract = new ethers.Contract(stockAddress, Stock.abi, signer)
      try {
        const data = await contract.sendStockTo(addressReceiver, amountStocksToSend)
        addressReceiver = 0
        amountStocksToSend = 0
      } catch (err) {
        console.log("Error: ", err)
      }
    }
  }

  async function shareDividentes() {
    if (typeof window.ethereum !== 'undefined') {
      await requestAccount()
      const provider = new ethers.providers.Web3Provider(window.ethereum)
      const signer = provider.getSigner()
      console.log({ provider })
      const contract = new ethers.Contract(stockAddress, Stock.abi, signer)
      try {
        await contract.distribute()
      } catch (err) {
        console.log("Error: ", err)
      }
    }
  }
</script>

<main>

	<div class="container">
		<div class="row justify-content-center">
			<h1>Decentralized {name}</h1>
		</div>
		<div class="row justify-content-center">
			<h4>Dapp for decentralized stock options</h4>
		</div>
		<div class="row justify-content-center">
			<span>Addresse: {stockAddress}</span>
		</div>

		<div class="row justify-content-center mt-5">
			<div class="col">
				<div class="card transfer">
				  <h5 class="card-header">Transfer stocks</h5>
				  <div class="card-body">
						<form>
						  <div class="form-group">
						    <label for="address">Address of receiver</label>
						    <input bind:value={addressReceiver} class="form-control" id="address"  placeholder="0">
								<label for="amount">Amount of stocks</label>
						    <input type="number" bind:value={amountStocksToSend} class="form-control" id="amount"  placeholder="0">
						  </div>
						</form>
						<Button on:click={transferStocks} class="btn btn-danger btn-lg active" role="button" aria-pressed="true">
							Transfer
						</Button>
				  </div>
				</div>
			</div>
			<div class="col">
				<div class="card share">
				  <h5 class="card-header">Share dividends</h5>
				  <div class="card-body">
				    <h5 class="card-title">Share dividents with all owners</h5>
				    <p class="card-text">the quantity of money on this address will be distributed to each holder of the stocks in proportion to the stocks he has.</p>
						<Button on:click={shareDividentes} class="btn btn-warning btn-lg active mt-5" role="button" aria-pressed="true">
							Share
						</Button>
				  </div>
				</div>
			</div>
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

	.transfer, .share {
		min-height: 320px;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
