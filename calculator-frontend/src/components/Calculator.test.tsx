import { render, screen, fireEvent } from '@testing-library/react';
import { describe, beforeEach, expect, test, vi } from 'vitest'
import '@testing-library/jest-dom';
import App from '../App';

describe('Calculator Component', () => {
  beforeEach(() => {
    render(<App />);
  });

  test('renders calculator component', () => {
    const calculatorElement = screen.getByTestId("calculator");
    expect(calculatorElement).toBeInTheDocument();
  });

  test('displays initial value 0', () => {
    const display = screen.getByDisplayValue('0');
    expect(display).toBeInTheDocument();
  });

  test('inputs numbers correctly', () => {
    const button7 = screen.getByText('7');
    fireEvent.click(button7);

    const display = screen.getByDisplayValue('7');
    expect(display).toBeInTheDocument();
  });

  test('handles addition operation', async () => {
    const button1 = screen.getByText('1');
    const button2 = screen.getByText('2');
    const buttonAdd = screen.getByText('+');
    const buttonEquals = screen.getByText('=');

    const mockResponse = { result: 13 }

    vi.spyOn(globalThis, 'fetch').mockResolvedValue({
      ok: true,
      json: async () => mockResponse,
    } as Response)

    /* fire events that update state */
    fireEvent.click(button1);
    fireEvent.click(button1);
    fireEvent.click(buttonAdd);
    fireEvent.click(button2);
    fireEvent.click(buttonEquals);

    const display = await screen.findByDisplayValue('13');
    expect(display).toBeInTheDocument();
  });

  test('handles square root operation', async () => {
    const button1 = screen.getByText('1');
    const button2 = screen.getByText('2');
    const buttonSqrt = screen.getByText('√');

    const mockResponse = { result: 11 }

    vi.spyOn(globalThis, 'fetch').mockResolvedValue({
      ok: true,
      json: async () => mockResponse,
    } as Response)

    fireEvent.click(button1);
    fireEvent.click(button2);
    fireEvent.click(button1);
    fireEvent.click(buttonSqrt);

    const display = await screen.findByDisplayValue('11');
    expect(display).toBeInTheDocument();
  });

  test('handles percentage operation', async () => {
    const button5 = screen.getByText('5');
    const button1 = screen.getByText('1');
    const buttonPercent = screen.getByText('%');
    const buttonEquals = screen.getByText('=');

    const mockResponse = { result: 140 }

    vi.spyOn(globalThis, 'fetch').mockResolvedValue({
      ok: true,
      json: async () => mockResponse,
    } as Response)

    fireEvent.click(button5);
    fireEvent.click(buttonPercent);
    fireEvent.click(button1);
    fireEvent.click(buttonEquals);

    const display = await screen.findByDisplayValue('140');
    expect(display).toBeInTheDocument();
  });

  test('handles clear operation', () => {
    const button1 = screen.getByText('1');
    const buttonClear = screen.getByText('AC');

    fireEvent.click(button1);
    fireEvent.click(buttonClear);

    const display = screen.getByDisplayValue('0');
    expect(display).toBeInTheDocument();
  });

  test('handles division by zero error', async () => {
    const button1 = screen.getByText('1');
    const button0 = screen.getByText('0');
    const buttonDivide = screen.getByText('÷');
    const buttonEquals = screen.getByText('=');

    // Mock fetch to return an error response for division by zero
    vi.spyOn(globalThis, 'fetch').mockResolvedValue({
      ok: false,
      json: async () => ({ error: "you can not divide by zero" }),
    } as Response)

    fireEvent.click(button1);
    fireEvent.click(buttonDivide);
    fireEvent.click(button0);
    fireEvent.click(buttonEquals);

    const toast = await screen.findByText('you can not divide by zero');

    expect(toast).toBeInTheDocument();
  });

  test('handles negative square root error', async () => {
    const buttonMinus = screen.getByText('-');
    const button1 = screen.getByText('1');
    const buttonSqrt = screen.getByText('√');

    // Mock fetch to return an error response for negative square root
    vi.spyOn(globalThis, 'fetch').mockResolvedValue({
      ok: false,
      json: async () => ({ error: "you can not get square root of a negative number" }),
    } as Response)

    fireEvent.click(buttonMinus);
    fireEvent.click(button1);
    fireEvent.click(buttonSqrt);

    const display = await screen.findByText('you can not get square root of a negative number');
    expect(display).toBeInTheDocument();
  });

  test('handles network error', async () => {
    const button1 = screen.getByText('1');
    const button2 = screen.getByText('2');
    const buttonAdd = screen.getByText('+');
    const buttonEquals = screen.getByText('=');

    // Mock fetch to return a network error
    vi.spyOn(globalThis, 'fetch').mockRejectedValue(new Error('Network error'));

    fireEvent.click(button1);
    fireEvent.click(buttonAdd);
    fireEvent.click(button2);
    fireEvent.click(buttonEquals);

    const display = await screen.findByText('Network error');
    expect(display).toBeInTheDocument();
  });

  test('handles decimal input correctly', () => {
    const button1 = screen.getByText('1');
    const buttonDot = screen.getByText('.');

    fireEvent.click(button1);
    fireEvent.click(buttonDot);
    fireEvent.click(button1);

    const display = screen.getByDisplayValue('1.1');
    expect(display).toBeInTheDocument();
  });

  test('handles toggle sign operation', () => {
    const button1 = screen.getByText('1');
    const buttonSign = screen.getByText('+/-');

    fireEvent.click(button1);
    fireEvent.click(buttonSign);

    const display = screen.getByDisplayValue('-1');
    expect(display).toBeInTheDocument();
  });

  test('handles multiple operations correctly', async () => {
    const button1 = screen.getByText('1');
    const button2 = screen.getByText('2');
    const buttonAdd = screen.getByText('+');
    const buttonEquals = screen.getByText('=');

    const mockResponse = { result: 13 }

    vi.spyOn(globalThis, 'fetch').mockResolvedValue({
      ok: true,
      json: async () => mockResponse,
    } as Response)

    fireEvent.click(button1);
    fireEvent.click(buttonAdd);
    fireEvent.click(button2);
    fireEvent.click(buttonEquals);

    // Check if result is displayed
    const display = await screen.findByDisplayValue('13');
    expect(display).toBeInTheDocument();
  });
});