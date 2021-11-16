const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Stock", function () {
  it("Should return the new balance after transfer", async function () {
    const Stock = await ethers.getContractFactory("Stock");
    const stock = await Stock.deploy();
    await stock.deployed();

    const balanceMinter = await stock.stockQte()
    expect(balanceMinter.toNumber() === 1000000)

    const balanceOneOwner = await stock.stockQteOf("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
    expect(balanceOneOwner.toNumber() === 0)
  });
});
