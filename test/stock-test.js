const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Stock", function () {
  it("Should return the new balance after transfer", async function () {
    const Stock = await ethers.getContractFactory("Stock");
    const stock = await Stock.deploy();
    await stock.deployed();

    const addressSecondOwner = "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"

    const balanceMinter = await stock.stockQte()
    expect(balanceMinter.toNumber() === 1000000)
    const nbrOwners = await stock.nbrOwners()
    expect(nbrOwners.toNumber() === 1)

    const balanceSecondOwner = await stock.stockQteOf(addressSecondOwner)
    expect(balanceSecondOwner.toNumber() === 0)

    await stock.sendStockTo(addressSecondOwner, 500000)
    const balanceSecondOwnerAfter = await stock.stockQteOf(addressSecondOwner)
    expect(balanceSecondOwnerAfter.toNumber() === 500000)

    const nbrOwnersAfter = await stock.nbrOwners()
    expect(nbrOwnersAfter.toNumber() === 2)

    const owners = await stock.getOwners()
    expect(owners.toString().split(',').length === 2)
  });
});
