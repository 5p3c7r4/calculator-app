import { useState } from "react";

// Calculator component for performing mathematical operations
function Calculator() {
  const [display, setDisplay] = useState("0");
  const [previousValue, setPreviousValue] = useState<number | null>(null);
  const [operation, setOperation] = useState<string | null>(null);
  const [waitingForOperand, setWaitingForOperand] = useState(false);

  const inputNumber = (num: string) => {
    if (waitingForOperand) {
      setDisplay(num);
      setWaitingForOperand(false);
    } else {
      setDisplay(display === "0" ? num : display + num);
    }
  };

  const inputDecimal = () => {
    if (waitingForOperand) {
      setDisplay("0.");
      setWaitingForOperand(false);
    } else if (display.indexOf(".") === -1) {
      setDisplay(display + ".");
    }
  };

  const clear = () => {
    setDisplay("0");
    setPreviousValue(null);
    setOperation(null);
    setWaitingForOperand(false);
  };

  const makeApiCall = async (operation: string, a: number, b?: number) => {
    try {
      const response = await fetch('api/v1/calculator', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          operation,
          a,
          b
        })
      });
      const result = await response.json();
      return result.result;
    } catch (error) {
      console.error('Error performing calculation:', error);
      throw new Error('Calculation failed');
    }
  };

  const performOperation = async (nextOperation: string) => {
    const inputValue = parseFloat(display);

    if (previousValue === null) {
      // First operand
      setPreviousValue(inputValue);
      setOperation(nextOperation);
      setWaitingForOperand(true);
    } else if (operation) {
      // Second operand - perform calculation
      const currentValue = previousValue || 0;

      try {
        const result = await makeApiCall(nextOperation, currentValue, inputValue);
        setDisplay(String(result));
        setPreviousValue(result);
        setOperation(null);
        setWaitingForOperand(true);
      } catch (error) {
        setDisplay('Error');
      }
    }
  };

  const handleEquals = async () => {
    if (previousValue !== null && operation) {
      try {
        const result = await makeApiCall(operation, previousValue, parseFloat(display));
        setDisplay(String(result));
        setPreviousValue(null);
        setOperation(null);
        setWaitingForOperand(true);
      } catch (error) {
        setDisplay('Error');
      }
    }
  };

  const handleSquareRoot = async () => {
    const inputValue = parseFloat(display);

    try {
      const result = await makeApiCall('sqrt', inputValue);
      setDisplay(String(result));
      setWaitingForOperand(true);
    } catch (error) {
      setDisplay('Error');
    }
  };

  const handleToggleSign = () => {
    if (display !== "0") {
      setDisplay(display.startsWith("-") ? display.slice(1) : "-" + display);
    }
  };

  return (
    <div className="app">
      <header className="header">
        <div className="header-content">
          <div className="brand">
            <div className="brand-icon">∑</div>
            Calculator
          </div>

          <span className="header-badge">Full-stack demo</span>
        </div>
      </header>

      <main className="main">
        <section className="hero">
          <span className="eyebrow">Arithmetic made simple</span>

          <h1>
            Calculate
            <br />
            <span>with confidence.</span>
          </h1>

          <p>
            A simple and intuitive calculator for basic and advanced
            mathematical operations.
          </p>

          <div className="features">
            <span>Basic operations</span>
            <span>Advanced arithmetic</span>
            <span>Responsive design</span>
          </div>
        </section>

        <div className="calculator">
          <div className="display">
            <input
              type="text"
              value={display}
              readOnly
              className="display-input"
            />
          </div>

          <div className="buttons">
            <button onClick={clear} className="btn clear">AC</button>
            <button onClick={handleToggleSign} className="btn sign">+/-</button>
            <button onClick={() => performOperation("percentage")} className="btn percentage">%</button>
            <button onClick={() => performOperation("divide")} className="btn operator">÷</button>

            <button onClick={() => inputNumber("7")} className="btn number">7</button>
            <button onClick={() => inputNumber("8")} className="btn number">8</button>
            <button onClick={() => inputNumber("9")} className="btn number">9</button>
            <button onClick={() => performOperation("multiply")} className="btn operator">×</button>

            <button onClick={() => inputNumber("4")} className="btn number">4</button>
            <button onClick={() => inputNumber("5")} className="btn number">5</button>
            <button onClick={() => inputNumber("6")} className="btn number">6</button>
            <button onClick={() => performOperation("subtract")} className="btn operator">−</button>

            <button onClick={() => inputNumber("1")} className="btn number">1</button>
            <button onClick={() => inputNumber("2")} className="btn number">2</button>
            <button onClick={() => inputNumber("3")} className="btn number">3</button>
            <button onClick={() => performOperation("add")} className="btn operator">+</button>

            <button onClick={() => inputNumber("0")} className="btn number zero">0</button>
            <button onClick={inputDecimal} className="btn number">.</button>
            <button onClick={handleEquals} className="btn equals">=</button>
            <button onClick={handleSquareRoot} className="btn advanced">√</button>
            <button onClick={() => performOperation("power")} className="btn advanced">xʸ</button>
          </div>
        </div>
      </main>

      <footer className="footer">
        Built with React + TypeScript
      </footer>
    </div>
  );
}

export default Calculator;
