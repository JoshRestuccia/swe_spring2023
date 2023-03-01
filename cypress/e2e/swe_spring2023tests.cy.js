describe("The Home Page", () => {
  it("successfully loads", () => {
    cy.visit("http://localhost:4200");
  });
});

describe("Button Test", () => {
  it('Clicks the "button", and checks for all dropdown options', () => {
    cy.visit("http://localhost:4200");

    cy.get("button").click();
    cy.contains("Login");
    cy.contains("Settings");
    cy.contains("Home");
  });
});

describe("Hyperlink Test", () => {
  it('clicking "Login" navigates to a new url', () => {
    cy.visit("http://localhost:4200");

    cy.get("button").click();
    cy.contains("Login").click();

    cy.url().should("include", "/login");
  });
});

describe("Box Test", () => {
  it("Tests Typing Into Boxes", () => {
    cy.visit("http://localhost:4200");

    cy.get("button").click();
    cy.contains("Login").click();

    cy.url().should("include", "/login");

    // Get an input, type into it
    cy.get("input").eq(0).type("email@com.com");
    cy.get("input").eq(0).should("have.value", "email@com.com");

    cy.get("input").eq(1).type("password");
    cy.get("input").eq(1).should("have.value", "password");
  });
});

describe("Hyperlink Test #2", () => {
  it('clicking "Home" navigates back to home', () => {
    cy.visit("http://localhost:4200");

    cy.get("button").click();
    cy.contains("Home").click();

    cy.url().should("include", "/mainpage");
  });
});

describe("Settings Button Check", () => {
  it("Checks the settings button can be pressed", () => {
    cy.visit("http://localhost:4200");

    cy.get("button").click();
    cy.contains("Settings").click();
  });
});
