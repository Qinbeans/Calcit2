// Generated from /Users/0x52696365/Documents/projects/Calcit2/calcit2.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class calcit2Parser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, GREATER=3, LESS=4, PLUS=5, MINUS=6, MULTIPLY=7, DIVIDE=8, 
		MODULO=9, NOT=10, LPAREN=11, RPAREN=12, COLON=13, EQUALS=14, POW=15, COLON_EQUALS=16, 
		SEMICOLON=17, BIT_AND=18, BIT_OR=19, BIT_XOR=20, BIT_CLEAR=21, LSHIFT=22, 
		RSHIFT=23, LOGICAL_AND=24, LOGICAL_OR=25, EQUAL=26, NOT_EQUAL=27, LESS_OR_EQUAL=28, 
		GREATER_OR_EQUAL=29, HEX_TYPE=30, BIN_TYPE=31, OCT_TYPE=32, F_TYPE=33, 
		I_TYPE=34, B_TYPE=35, S_TYPE=36, P_CAST=37, FUNCTION_NAME=38, IDENTIFIER=39, 
		HEX_LITERAL=40, BIN_LITERAL=41, OCT_LITERAL=42, INTEGER_LITERAL=43, FLOAT_LITERAL=44, 
		CHAR_LITERAL=45, WS=46;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_function_assignment = 2, RULE_expression = 3, 
		RULE_number = 4, RULE_any = 5, RULE_integer_expression = 6, RULE_integer_cast_expression = 7, 
		RULE_integer_pcast_expression = 8, RULE_float_expression = 9, RULE_float_cast_expression = 10, 
		RULE_float_pcast_expression = 11, RULE_bool_expression = 12, RULE_bool_cast_expression = 13, 
		RULE_bool_number_operation = 14;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "function_assignment", "expression", "number", 
			"any", "integer_expression", "integer_cast_expression", "integer_pcast_expression", 
			"float_expression", "float_cast_expression", "float_pcast_expression", 
			"bool_expression", "bool_cast_expression", "bool_number_operation"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'true'", "'false'", "'>'", "'<'", "'+'", "'-'", "'*'", "'/'", 
			"'%'", "'!'", "'('", "')'", "':'", "'='", "'**'", "':='", "';'", "'&'", 
			"'|'", "'^'", "'&^'", "'<<'", "'>>'", "'&&'", "'||'", "'=='", "'!='", 
			"'<='", "'>='", "'hex'", "'bin'", "'oct'", "'float'", "'int'", "'bool'", 
			"'string'", "'pcast'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, "GREATER", "LESS", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", 
			"MODULO", "NOT", "LPAREN", "RPAREN", "COLON", "EQUALS", "POW", "COLON_EQUALS", 
			"SEMICOLON", "BIT_AND", "BIT_OR", "BIT_XOR", "BIT_CLEAR", "LSHIFT", "RSHIFT", 
			"LOGICAL_AND", "LOGICAL_OR", "EQUAL", "NOT_EQUAL", "LESS_OR_EQUAL", "GREATER_OR_EQUAL", 
			"HEX_TYPE", "BIN_TYPE", "OCT_TYPE", "F_TYPE", "I_TYPE", "B_TYPE", "S_TYPE", 
			"P_CAST", "FUNCTION_NAME", "IDENTIFIER", "HEX_LITERAL", "BIN_LITERAL", 
			"OCT_LITERAL", "INTEGER_LITERAL", "FLOAT_LITERAL", "CHAR_LITERAL", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "calcit2.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public calcit2Parser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(33);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 35114578873414L) != 0)) {
				{
				{
				setState(30);
				statement();
				}
				}
				setState(35);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public Function_assignmentContext function_assignment() {
			return getRuleContext(Function_assignmentContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(38);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FUNCTION_NAME:
				enterOuterAlt(_localctx, 1);
				{
				setState(36);
				function_assignment();
				}
				break;
			case T__0:
			case T__1:
			case MINUS:
			case NOT:
			case LPAREN:
			case HEX_TYPE:
			case BIN_TYPE:
			case OCT_TYPE:
			case F_TYPE:
			case I_TYPE:
			case B_TYPE:
			case P_CAST:
			case IDENTIFIER:
			case HEX_LITERAL:
			case BIN_LITERAL:
			case OCT_LITERAL:
			case INTEGER_LITERAL:
			case FLOAT_LITERAL:
				enterOuterAlt(_localctx, 2);
				{
				setState(37);
				expression();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Function_assignmentContext extends ParserRuleContext {
		public TerminalNode FUNCTION_NAME() { return getToken(calcit2Parser.FUNCTION_NAME, 0); }
		public TerminalNode LPAREN() { return getToken(calcit2Parser.LPAREN, 0); }
		public TerminalNode IDENTIFIER() { return getToken(calcit2Parser.IDENTIFIER, 0); }
		public TerminalNode RPAREN() { return getToken(calcit2Parser.RPAREN, 0); }
		public TerminalNode EQUALS() { return getToken(calcit2Parser.EQUALS, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Function_assignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_assignment; }
	}

	public final Function_assignmentContext function_assignment() throws RecognitionException {
		Function_assignmentContext _localctx = new Function_assignmentContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_function_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(40);
			match(FUNCTION_NAME);
			setState(41);
			match(LPAREN);
			setState(42);
			match(IDENTIFIER);
			setState(43);
			match(RPAREN);
			setState(44);
			match(EQUALS);
			setState(45);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public Integer_expressionContext integer_expression() {
			return getRuleContext(Integer_expressionContext.class,0);
		}
		public Float_expressionContext float_expression() {
			return getRuleContext(Float_expressionContext.class,0);
		}
		public Bool_expressionContext bool_expression() {
			return getRuleContext(Bool_expressionContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_expression);
		try {
			setState(50);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(47);
				integer_expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(48);
				float_expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(49);
				bool_expression(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class NumberContext extends ParserRuleContext {
		public Integer_expressionContext integer_expression() {
			return getRuleContext(Integer_expressionContext.class,0);
		}
		public Float_expressionContext float_expression() {
			return getRuleContext(Float_expressionContext.class,0);
		}
		public NumberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_number; }
	}

	public final NumberContext number() throws RecognitionException {
		NumberContext _localctx = new NumberContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_number);
		try {
			setState(54);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(52);
				integer_expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(53);
				float_expression(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AnyContext extends ParserRuleContext {
		public NumberContext number() {
			return getRuleContext(NumberContext.class,0);
		}
		public Bool_expressionContext bool_expression() {
			return getRuleContext(Bool_expressionContext.class,0);
		}
		public AnyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_any; }
	}

	public final AnyContext any() throws RecognitionException {
		AnyContext _localctx = new AnyContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_any);
		try {
			setState(58);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(56);
				number();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(57);
				bool_expression(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Integer_expressionContext extends ParserRuleContext {
		public Integer_expressionContext left;
		public Integer_expressionContext right;
		public Token op;
		public TerminalNode LPAREN() { return getToken(calcit2Parser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(calcit2Parser.RPAREN, 0); }
		public List<Integer_expressionContext> integer_expression() {
			return getRuleContexts(Integer_expressionContext.class);
		}
		public Integer_expressionContext integer_expression(int i) {
			return getRuleContext(Integer_expressionContext.class,i);
		}
		public Integer_cast_expressionContext integer_cast_expression() {
			return getRuleContext(Integer_cast_expressionContext.class,0);
		}
		public Integer_pcast_expressionContext integer_pcast_expression() {
			return getRuleContext(Integer_pcast_expressionContext.class,0);
		}
		public TerminalNode MINUS() { return getToken(calcit2Parser.MINUS, 0); }
		public TerminalNode NOT() { return getToken(calcit2Parser.NOT, 0); }
		public TerminalNode INTEGER_LITERAL() { return getToken(calcit2Parser.INTEGER_LITERAL, 0); }
		public TerminalNode HEX_LITERAL() { return getToken(calcit2Parser.HEX_LITERAL, 0); }
		public TerminalNode BIN_LITERAL() { return getToken(calcit2Parser.BIN_LITERAL, 0); }
		public TerminalNode OCT_LITERAL() { return getToken(calcit2Parser.OCT_LITERAL, 0); }
		public TerminalNode IDENTIFIER() { return getToken(calcit2Parser.IDENTIFIER, 0); }
		public TerminalNode POW() { return getToken(calcit2Parser.POW, 0); }
		public TerminalNode MULTIPLY() { return getToken(calcit2Parser.MULTIPLY, 0); }
		public TerminalNode DIVIDE() { return getToken(calcit2Parser.DIVIDE, 0); }
		public TerminalNode MODULO() { return getToken(calcit2Parser.MODULO, 0); }
		public TerminalNode PLUS() { return getToken(calcit2Parser.PLUS, 0); }
		public TerminalNode BIT_AND() { return getToken(calcit2Parser.BIT_AND, 0); }
		public TerminalNode BIT_OR() { return getToken(calcit2Parser.BIT_OR, 0); }
		public TerminalNode BIT_XOR() { return getToken(calcit2Parser.BIT_XOR, 0); }
		public TerminalNode BIT_CLEAR() { return getToken(calcit2Parser.BIT_CLEAR, 0); }
		public TerminalNode LSHIFT() { return getToken(calcit2Parser.LSHIFT, 0); }
		public TerminalNode RSHIFT() { return getToken(calcit2Parser.RSHIFT, 0); }
		public Integer_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_integer_expression; }
	}

	public final Integer_expressionContext integer_expression() throws RecognitionException {
		return integer_expression(0);
	}

	private Integer_expressionContext integer_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Integer_expressionContext _localctx = new Integer_expressionContext(_ctx, _parentState);
		Integer_expressionContext _prevctx = _localctx;
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_integer_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(72);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LPAREN:
				{
				setState(61);
				match(LPAREN);
				setState(62);
				((Integer_expressionContext)_localctx).right = integer_expression(0);
				setState(63);
				match(RPAREN);
				}
				break;
			case HEX_TYPE:
			case BIN_TYPE:
			case OCT_TYPE:
			case F_TYPE:
			case I_TYPE:
				{
				setState(65);
				integer_cast_expression();
				}
				break;
			case P_CAST:
				{
				setState(66);
				integer_pcast_expression();
				}
				break;
			case MINUS:
				{
				setState(67);
				((Integer_expressionContext)_localctx).op = match(MINUS);
				setState(68);
				((Integer_expressionContext)_localctx).right = integer_expression(7);
				}
				break;
			case NOT:
				{
				setState(69);
				((Integer_expressionContext)_localctx).op = match(NOT);
				setState(70);
				((Integer_expressionContext)_localctx).right = integer_expression(6);
				}
				break;
			case IDENTIFIER:
			case HEX_LITERAL:
			case BIN_LITERAL:
			case OCT_LITERAL:
			case INTEGER_LITERAL:
				{
				setState(71);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 17042430230528L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(88);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(86);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
					case 1:
						{
						_localctx = new Integer_expressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_integer_expression);
						setState(74);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(75);
						((Integer_expressionContext)_localctx).op = match(POW);
						setState(76);
						((Integer_expressionContext)_localctx).right = integer_expression(6);
						}
						break;
					case 2:
						{
						_localctx = new Integer_expressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_integer_expression);
						setState(77);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(78);
						((Integer_expressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 896L) != 0)) ) {
							((Integer_expressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(79);
						((Integer_expressionContext)_localctx).right = integer_expression(5);
						}
						break;
					case 3:
						{
						_localctx = new Integer_expressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_integer_expression);
						setState(80);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(81);
						((Integer_expressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==PLUS || _la==MINUS) ) {
							((Integer_expressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(82);
						((Integer_expressionContext)_localctx).right = integer_expression(4);
						}
						break;
					case 4:
						{
						_localctx = new Integer_expressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_integer_expression);
						setState(83);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(84);
						((Integer_expressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 16515072L) != 0)) ) {
							((Integer_expressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(85);
						((Integer_expressionContext)_localctx).right = integer_expression(3);
						}
						break;
					}
					} 
				}
				setState(90);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Integer_cast_expressionContext extends ParserRuleContext {
		public Token op;
		public NumberContext right;
		public TerminalNode LPAREN() { return getToken(calcit2Parser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(calcit2Parser.RPAREN, 0); }
		public NumberContext number() {
			return getRuleContext(NumberContext.class,0);
		}
		public TerminalNode HEX_TYPE() { return getToken(calcit2Parser.HEX_TYPE, 0); }
		public TerminalNode BIN_TYPE() { return getToken(calcit2Parser.BIN_TYPE, 0); }
		public TerminalNode OCT_TYPE() { return getToken(calcit2Parser.OCT_TYPE, 0); }
		public TerminalNode I_TYPE() { return getToken(calcit2Parser.I_TYPE, 0); }
		public TerminalNode F_TYPE() { return getToken(calcit2Parser.F_TYPE, 0); }
		public Integer_cast_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_integer_cast_expression; }
	}

	public final Integer_cast_expressionContext integer_cast_expression() throws RecognitionException {
		Integer_cast_expressionContext _localctx = new Integer_cast_expressionContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_integer_cast_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(91);
			((Integer_cast_expressionContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 33285996544L) != 0)) ) {
				((Integer_cast_expressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(92);
			match(LPAREN);
			setState(93);
			((Integer_cast_expressionContext)_localctx).right = number();
			setState(94);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Integer_pcast_expressionContext extends ParserRuleContext {
		public Token op;
		public Float_expressionContext right;
		public TerminalNode LPAREN() { return getToken(calcit2Parser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(calcit2Parser.RPAREN, 0); }
		public TerminalNode P_CAST() { return getToken(calcit2Parser.P_CAST, 0); }
		public Float_expressionContext float_expression() {
			return getRuleContext(Float_expressionContext.class,0);
		}
		public Integer_pcast_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_integer_pcast_expression; }
	}

	public final Integer_pcast_expressionContext integer_pcast_expression() throws RecognitionException {
		Integer_pcast_expressionContext _localctx = new Integer_pcast_expressionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_integer_pcast_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(96);
			((Integer_pcast_expressionContext)_localctx).op = match(P_CAST);
			setState(97);
			match(LPAREN);
			setState(98);
			((Integer_pcast_expressionContext)_localctx).right = float_expression(0);
			setState(99);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Float_expressionContext extends ParserRuleContext {
		public Float_expressionContext left;
		public Float_expressionContext right;
		public Token op;
		public TerminalNode LPAREN() { return getToken(calcit2Parser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(calcit2Parser.RPAREN, 0); }
		public List<Float_expressionContext> float_expression() {
			return getRuleContexts(Float_expressionContext.class);
		}
		public Float_expressionContext float_expression(int i) {
			return getRuleContext(Float_expressionContext.class,i);
		}
		public Float_cast_expressionContext float_cast_expression() {
			return getRuleContext(Float_cast_expressionContext.class,0);
		}
		public Float_pcast_expressionContext float_pcast_expression() {
			return getRuleContext(Float_pcast_expressionContext.class,0);
		}
		public TerminalNode MINUS() { return getToken(calcit2Parser.MINUS, 0); }
		public TerminalNode FLOAT_LITERAL() { return getToken(calcit2Parser.FLOAT_LITERAL, 0); }
		public TerminalNode IDENTIFIER() { return getToken(calcit2Parser.IDENTIFIER, 0); }
		public TerminalNode POW() { return getToken(calcit2Parser.POW, 0); }
		public TerminalNode MULTIPLY() { return getToken(calcit2Parser.MULTIPLY, 0); }
		public TerminalNode DIVIDE() { return getToken(calcit2Parser.DIVIDE, 0); }
		public TerminalNode PLUS() { return getToken(calcit2Parser.PLUS, 0); }
		public Float_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_float_expression; }
	}

	public final Float_expressionContext float_expression() throws RecognitionException {
		return float_expression(0);
	}

	private Float_expressionContext float_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Float_expressionContext _localctx = new Float_expressionContext(_ctx, _parentState);
		Float_expressionContext _prevctx = _localctx;
		int _startState = 18;
		enterRecursionRule(_localctx, 18, RULE_float_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(111);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LPAREN:
				{
				setState(102);
				match(LPAREN);
				setState(103);
				((Float_expressionContext)_localctx).right = float_expression(0);
				setState(104);
				match(RPAREN);
				}
				break;
			case F_TYPE:
				{
				setState(106);
				float_cast_expression();
				}
				break;
			case P_CAST:
				{
				setState(107);
				float_pcast_expression();
				}
				break;
			case MINUS:
				{
				setState(108);
				((Float_expressionContext)_localctx).op = match(MINUS);
				setState(109);
				((Float_expressionContext)_localctx).right = float_expression(5);
				}
				break;
			case IDENTIFIER:
			case FLOAT_LITERAL:
				{
				setState(110);
				_la = _input.LA(1);
				if ( !(_la==IDENTIFIER || _la==FLOAT_LITERAL) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(124);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(122);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
					case 1:
						{
						_localctx = new Float_expressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_float_expression);
						setState(113);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(114);
						((Float_expressionContext)_localctx).op = match(POW);
						setState(115);
						((Float_expressionContext)_localctx).right = float_expression(5);
						}
						break;
					case 2:
						{
						_localctx = new Float_expressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_float_expression);
						setState(116);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(117);
						((Float_expressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MULTIPLY || _la==DIVIDE) ) {
							((Float_expressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(118);
						((Float_expressionContext)_localctx).right = float_expression(4);
						}
						break;
					case 3:
						{
						_localctx = new Float_expressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_float_expression);
						setState(119);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(120);
						((Float_expressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==PLUS || _la==MINUS) ) {
							((Float_expressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(121);
						((Float_expressionContext)_localctx).right = float_expression(3);
						}
						break;
					}
					} 
				}
				setState(126);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Float_cast_expressionContext extends ParserRuleContext {
		public Token op;
		public NumberContext right;
		public TerminalNode LPAREN() { return getToken(calcit2Parser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(calcit2Parser.RPAREN, 0); }
		public TerminalNode F_TYPE() { return getToken(calcit2Parser.F_TYPE, 0); }
		public NumberContext number() {
			return getRuleContext(NumberContext.class,0);
		}
		public Float_cast_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_float_cast_expression; }
	}

	public final Float_cast_expressionContext float_cast_expression() throws RecognitionException {
		Float_cast_expressionContext _localctx = new Float_cast_expressionContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_float_cast_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(127);
			((Float_cast_expressionContext)_localctx).op = match(F_TYPE);
			setState(128);
			match(LPAREN);
			setState(129);
			((Float_cast_expressionContext)_localctx).right = number();
			setState(130);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Float_pcast_expressionContext extends ParserRuleContext {
		public Token op;
		public Integer_expressionContext right;
		public TerminalNode LPAREN() { return getToken(calcit2Parser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(calcit2Parser.RPAREN, 0); }
		public TerminalNode P_CAST() { return getToken(calcit2Parser.P_CAST, 0); }
		public Integer_expressionContext integer_expression() {
			return getRuleContext(Integer_expressionContext.class,0);
		}
		public Float_pcast_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_float_pcast_expression; }
	}

	public final Float_pcast_expressionContext float_pcast_expression() throws RecognitionException {
		Float_pcast_expressionContext _localctx = new Float_pcast_expressionContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_float_pcast_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(132);
			((Float_pcast_expressionContext)_localctx).op = match(P_CAST);
			setState(133);
			match(LPAREN);
			setState(134);
			((Float_pcast_expressionContext)_localctx).right = integer_expression(0);
			setState(135);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Bool_expressionContext extends ParserRuleContext {
		public Bool_expressionContext left;
		public Bool_expressionContext right;
		public Token op;
		public TerminalNode LPAREN() { return getToken(calcit2Parser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(calcit2Parser.RPAREN, 0); }
		public List<Bool_expressionContext> bool_expression() {
			return getRuleContexts(Bool_expressionContext.class);
		}
		public Bool_expressionContext bool_expression(int i) {
			return getRuleContext(Bool_expressionContext.class,i);
		}
		public Bool_cast_expressionContext bool_cast_expression() {
			return getRuleContext(Bool_cast_expressionContext.class,0);
		}
		public Bool_number_operationContext bool_number_operation() {
			return getRuleContext(Bool_number_operationContext.class,0);
		}
		public TerminalNode NOT() { return getToken(calcit2Parser.NOT, 0); }
		public TerminalNode EQUAL() { return getToken(calcit2Parser.EQUAL, 0); }
		public TerminalNode NOT_EQUAL() { return getToken(calcit2Parser.NOT_EQUAL, 0); }
		public TerminalNode LOGICAL_AND() { return getToken(calcit2Parser.LOGICAL_AND, 0); }
		public TerminalNode LOGICAL_OR() { return getToken(calcit2Parser.LOGICAL_OR, 0); }
		public Bool_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bool_expression; }
	}

	public final Bool_expressionContext bool_expression() throws RecognitionException {
		return bool_expression(0);
	}

	private Bool_expressionContext bool_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Bool_expressionContext _localctx = new Bool_expressionContext(_ctx, _parentState);
		Bool_expressionContext _prevctx = _localctx;
		int _startState = 24;
		enterRecursionRule(_localctx, 24, RULE_bool_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				setState(138);
				match(LPAREN);
				setState(139);
				((Bool_expressionContext)_localctx).right = bool_expression(0);
				setState(140);
				match(RPAREN);
				}
				break;
			case 2:
				{
				setState(142);
				bool_cast_expression();
				}
				break;
			case 3:
				{
				setState(143);
				bool_number_operation();
				}
				break;
			case 4:
				{
				setState(144);
				((Bool_expressionContext)_localctx).op = match(NOT);
				setState(145);
				((Bool_expressionContext)_localctx).right = bool_expression(2);
				}
				break;
			case 5:
				{
				setState(146);
				_la = _input.LA(1);
				if ( !(_la==T__0 || _la==T__1) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(157);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(155);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
					case 1:
						{
						_localctx = new Bool_expressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_bool_expression);
						setState(149);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(150);
						((Bool_expressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==EQUAL || _la==NOT_EQUAL) ) {
							((Bool_expressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(151);
						((Bool_expressionContext)_localctx).right = bool_expression(5);
						}
						break;
					case 2:
						{
						_localctx = new Bool_expressionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_bool_expression);
						setState(152);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(153);
						((Bool_expressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==LOGICAL_AND || _la==LOGICAL_OR) ) {
							((Bool_expressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(154);
						((Bool_expressionContext)_localctx).right = bool_expression(4);
						}
						break;
					}
					} 
				}
				setState(159);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Bool_cast_expressionContext extends ParserRuleContext {
		public Token op;
		public AnyContext right;
		public TerminalNode LPAREN() { return getToken(calcit2Parser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(calcit2Parser.RPAREN, 0); }
		public TerminalNode B_TYPE() { return getToken(calcit2Parser.B_TYPE, 0); }
		public AnyContext any() {
			return getRuleContext(AnyContext.class,0);
		}
		public Bool_cast_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bool_cast_expression; }
	}

	public final Bool_cast_expressionContext bool_cast_expression() throws RecognitionException {
		Bool_cast_expressionContext _localctx = new Bool_cast_expressionContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_bool_cast_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(160);
			((Bool_cast_expressionContext)_localctx).op = match(B_TYPE);
			setState(161);
			match(LPAREN);
			setState(162);
			((Bool_cast_expressionContext)_localctx).right = any();
			setState(163);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Bool_number_operationContext extends ParserRuleContext {
		public NumberContext left;
		public Token op;
		public NumberContext right;
		public List<NumberContext> number() {
			return getRuleContexts(NumberContext.class);
		}
		public NumberContext number(int i) {
			return getRuleContext(NumberContext.class,i);
		}
		public TerminalNode EQUAL() { return getToken(calcit2Parser.EQUAL, 0); }
		public TerminalNode NOT_EQUAL() { return getToken(calcit2Parser.NOT_EQUAL, 0); }
		public TerminalNode LESS() { return getToken(calcit2Parser.LESS, 0); }
		public TerminalNode GREATER() { return getToken(calcit2Parser.GREATER, 0); }
		public TerminalNode LESS_OR_EQUAL() { return getToken(calcit2Parser.LESS_OR_EQUAL, 0); }
		public TerminalNode GREATER_OR_EQUAL() { return getToken(calcit2Parser.GREATER_OR_EQUAL, 0); }
		public Bool_number_operationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bool_number_operation; }
	}

	public final Bool_number_operationContext bool_number_operation() throws RecognitionException {
		Bool_number_operationContext _localctx = new Bool_number_operationContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_bool_number_operation);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(165);
			((Bool_number_operationContext)_localctx).left = number();
			setState(166);
			((Bool_number_operationContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1006632984L) != 0)) ) {
				((Bool_number_operationContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(167);
			((Bool_number_operationContext)_localctx).right = number();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 6:
			return integer_expression_sempred((Integer_expressionContext)_localctx, predIndex);
		case 9:
			return float_expression_sempred((Float_expressionContext)_localctx, predIndex);
		case 12:
			return bool_expression_sempred((Bool_expressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean integer_expression_sempred(Integer_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 5);
		case 1:
			return precpred(_ctx, 4);
		case 2:
			return precpred(_ctx, 3);
		case 3:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean float_expression_sempred(Float_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 4);
		case 5:
			return precpred(_ctx, 3);
		case 6:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean bool_expression_sempred(Bool_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 7:
			return precpred(_ctx, 4);
		case 8:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001.\u00aa\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0001\u0000\u0005\u0000"+
		" \b\u0000\n\u0000\f\u0000#\t\u0000\u0001\u0001\u0001\u0001\u0003\u0001"+
		"\'\b\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003"+
		"3\b\u0003\u0001\u0004\u0001\u0004\u0003\u00047\b\u0004\u0001\u0005\u0001"+
		"\u0005\u0003\u0005;\b\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0003\u0006I\b\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0005\u0006W\b"+
		"\u0006\n\u0006\f\u0006Z\t\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0003"+
		"\tp\b\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0005\t{\b\t\n\t\f\t~\t\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\f"+
		"\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0003\f\u0094\b\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0005"+
		"\f\u009c\b\f\n\f\f\f\u009f\t\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0000\u0003"+
		"\f\u0012\u0018\u000f\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014"+
		"\u0016\u0018\u001a\u001c\u0000\u000b\u0001\u0000\'+\u0001\u0000\u0007"+
		"\t\u0001\u0000\u0005\u0006\u0001\u0000\u0012\u0017\u0001\u0000\u001e\""+
		"\u0002\u0000\'\',,\u0001\u0000\u0007\b\u0001\u0000\u0001\u0002\u0001\u0000"+
		"\u001a\u001b\u0001\u0000\u0018\u0019\u0002\u0000\u0003\u0004\u001a\u001d"+
		"\u00b6\u0000!\u0001\u0000\u0000\u0000\u0002&\u0001\u0000\u0000\u0000\u0004"+
		"(\u0001\u0000\u0000\u0000\u00062\u0001\u0000\u0000\u0000\b6\u0001\u0000"+
		"\u0000\u0000\n:\u0001\u0000\u0000\u0000\fH\u0001\u0000\u0000\u0000\u000e"+
		"[\u0001\u0000\u0000\u0000\u0010`\u0001\u0000\u0000\u0000\u0012o\u0001"+
		"\u0000\u0000\u0000\u0014\u007f\u0001\u0000\u0000\u0000\u0016\u0084\u0001"+
		"\u0000\u0000\u0000\u0018\u0093\u0001\u0000\u0000\u0000\u001a\u00a0\u0001"+
		"\u0000\u0000\u0000\u001c\u00a5\u0001\u0000\u0000\u0000\u001e \u0003\u0002"+
		"\u0001\u0000\u001f\u001e\u0001\u0000\u0000\u0000 #\u0001\u0000\u0000\u0000"+
		"!\u001f\u0001\u0000\u0000\u0000!\"\u0001\u0000\u0000\u0000\"\u0001\u0001"+
		"\u0000\u0000\u0000#!\u0001\u0000\u0000\u0000$\'\u0003\u0004\u0002\u0000"+
		"%\'\u0003\u0006\u0003\u0000&$\u0001\u0000\u0000\u0000&%\u0001\u0000\u0000"+
		"\u0000\'\u0003\u0001\u0000\u0000\u0000()\u0005&\u0000\u0000)*\u0005\u000b"+
		"\u0000\u0000*+\u0005\'\u0000\u0000+,\u0005\f\u0000\u0000,-\u0005\u000e"+
		"\u0000\u0000-.\u0003\u0006\u0003\u0000.\u0005\u0001\u0000\u0000\u0000"+
		"/3\u0003\f\u0006\u000003\u0003\u0012\t\u000013\u0003\u0018\f\u00002/\u0001"+
		"\u0000\u0000\u000020\u0001\u0000\u0000\u000021\u0001\u0000\u0000\u0000"+
		"3\u0007\u0001\u0000\u0000\u000047\u0003\f\u0006\u000057\u0003\u0012\t"+
		"\u000064\u0001\u0000\u0000\u000065\u0001\u0000\u0000\u00007\t\u0001\u0000"+
		"\u0000\u00008;\u0003\b\u0004\u00009;\u0003\u0018\f\u0000:8\u0001\u0000"+
		"\u0000\u0000:9\u0001\u0000\u0000\u0000;\u000b\u0001\u0000\u0000\u0000"+
		"<=\u0006\u0006\uffff\uffff\u0000=>\u0005\u000b\u0000\u0000>?\u0003\f\u0006"+
		"\u0000?@\u0005\f\u0000\u0000@I\u0001\u0000\u0000\u0000AI\u0003\u000e\u0007"+
		"\u0000BI\u0003\u0010\b\u0000CD\u0005\u0006\u0000\u0000DI\u0003\f\u0006"+
		"\u0007EF\u0005\n\u0000\u0000FI\u0003\f\u0006\u0006GI\u0007\u0000\u0000"+
		"\u0000H<\u0001\u0000\u0000\u0000HA\u0001\u0000\u0000\u0000HB\u0001\u0000"+
		"\u0000\u0000HC\u0001\u0000\u0000\u0000HE\u0001\u0000\u0000\u0000HG\u0001"+
		"\u0000\u0000\u0000IX\u0001\u0000\u0000\u0000JK\n\u0005\u0000\u0000KL\u0005"+
		"\u000f\u0000\u0000LW\u0003\f\u0006\u0006MN\n\u0004\u0000\u0000NO\u0007"+
		"\u0001\u0000\u0000OW\u0003\f\u0006\u0005PQ\n\u0003\u0000\u0000QR\u0007"+
		"\u0002\u0000\u0000RW\u0003\f\u0006\u0004ST\n\u0002\u0000\u0000TU\u0007"+
		"\u0003\u0000\u0000UW\u0003\f\u0006\u0003VJ\u0001\u0000\u0000\u0000VM\u0001"+
		"\u0000\u0000\u0000VP\u0001\u0000\u0000\u0000VS\u0001\u0000\u0000\u0000"+
		"WZ\u0001\u0000\u0000\u0000XV\u0001\u0000\u0000\u0000XY\u0001\u0000\u0000"+
		"\u0000Y\r\u0001\u0000\u0000\u0000ZX\u0001\u0000\u0000\u0000[\\\u0007\u0004"+
		"\u0000\u0000\\]\u0005\u000b\u0000\u0000]^\u0003\b\u0004\u0000^_\u0005"+
		"\f\u0000\u0000_\u000f\u0001\u0000\u0000\u0000`a\u0005%\u0000\u0000ab\u0005"+
		"\u000b\u0000\u0000bc\u0003\u0012\t\u0000cd\u0005\f\u0000\u0000d\u0011"+
		"\u0001\u0000\u0000\u0000ef\u0006\t\uffff\uffff\u0000fg\u0005\u000b\u0000"+
		"\u0000gh\u0003\u0012\t\u0000hi\u0005\f\u0000\u0000ip\u0001\u0000\u0000"+
		"\u0000jp\u0003\u0014\n\u0000kp\u0003\u0016\u000b\u0000lm\u0005\u0006\u0000"+
		"\u0000mp\u0003\u0012\t\u0005np\u0007\u0005\u0000\u0000oe\u0001\u0000\u0000"+
		"\u0000oj\u0001\u0000\u0000\u0000ok\u0001\u0000\u0000\u0000ol\u0001\u0000"+
		"\u0000\u0000on\u0001\u0000\u0000\u0000p|\u0001\u0000\u0000\u0000qr\n\u0004"+
		"\u0000\u0000rs\u0005\u000f\u0000\u0000s{\u0003\u0012\t\u0005tu\n\u0003"+
		"\u0000\u0000uv\u0007\u0006\u0000\u0000v{\u0003\u0012\t\u0004wx\n\u0002"+
		"\u0000\u0000xy\u0007\u0002\u0000\u0000y{\u0003\u0012\t\u0003zq\u0001\u0000"+
		"\u0000\u0000zt\u0001\u0000\u0000\u0000zw\u0001\u0000\u0000\u0000{~\u0001"+
		"\u0000\u0000\u0000|z\u0001\u0000\u0000\u0000|}\u0001\u0000\u0000\u0000"+
		"}\u0013\u0001\u0000\u0000\u0000~|\u0001\u0000\u0000\u0000\u007f\u0080"+
		"\u0005!\u0000\u0000\u0080\u0081\u0005\u000b\u0000\u0000\u0081\u0082\u0003"+
		"\b\u0004\u0000\u0082\u0083\u0005\f\u0000\u0000\u0083\u0015\u0001\u0000"+
		"\u0000\u0000\u0084\u0085\u0005%\u0000\u0000\u0085\u0086\u0005\u000b\u0000"+
		"\u0000\u0086\u0087\u0003\f\u0006\u0000\u0087\u0088\u0005\f\u0000\u0000"+
		"\u0088\u0017\u0001\u0000\u0000\u0000\u0089\u008a\u0006\f\uffff\uffff\u0000"+
		"\u008a\u008b\u0005\u000b\u0000\u0000\u008b\u008c\u0003\u0018\f\u0000\u008c"+
		"\u008d\u0005\f\u0000\u0000\u008d\u0094\u0001\u0000\u0000\u0000\u008e\u0094"+
		"\u0003\u001a\r\u0000\u008f\u0094\u0003\u001c\u000e\u0000\u0090\u0091\u0005"+
		"\n\u0000\u0000\u0091\u0094\u0003\u0018\f\u0002\u0092\u0094\u0007\u0007"+
		"\u0000\u0000\u0093\u0089\u0001\u0000\u0000\u0000\u0093\u008e\u0001\u0000"+
		"\u0000\u0000\u0093\u008f\u0001\u0000\u0000\u0000\u0093\u0090\u0001\u0000"+
		"\u0000\u0000\u0093\u0092\u0001\u0000\u0000\u0000\u0094\u009d\u0001\u0000"+
		"\u0000\u0000\u0095\u0096\n\u0004\u0000\u0000\u0096\u0097\u0007\b\u0000"+
		"\u0000\u0097\u009c\u0003\u0018\f\u0005\u0098\u0099\n\u0003\u0000\u0000"+
		"\u0099\u009a\u0007\t\u0000\u0000\u009a\u009c\u0003\u0018\f\u0004\u009b"+
		"\u0095\u0001\u0000\u0000\u0000\u009b\u0098\u0001\u0000\u0000\u0000\u009c"+
		"\u009f\u0001\u0000\u0000\u0000\u009d\u009b\u0001\u0000\u0000\u0000\u009d"+
		"\u009e\u0001\u0000\u0000\u0000\u009e\u0019\u0001\u0000\u0000\u0000\u009f"+
		"\u009d\u0001\u0000\u0000\u0000\u00a0\u00a1\u0005#\u0000\u0000\u00a1\u00a2"+
		"\u0005\u000b\u0000\u0000\u00a2\u00a3\u0003\n\u0005\u0000\u00a3\u00a4\u0005"+
		"\f\u0000\u0000\u00a4\u001b\u0001\u0000\u0000\u0000\u00a5\u00a6\u0003\b"+
		"\u0004\u0000\u00a6\u00a7\u0007\n\u0000\u0000\u00a7\u00a8\u0003\b\u0004"+
		"\u0000\u00a8\u001d\u0001\u0000\u0000\u0000\u000e!&26:HVXoz|\u0093\u009b"+
		"\u009d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}