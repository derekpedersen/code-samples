using System;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using CodeFights.Arcade;

namespace UnitTestProject1.ArcadeTests
{
    [TestClass]
    public class SmoothSailingTests
    {
        private SmoothSailing _service { get; set; }

        public SmoothSailingTests()
        {
            this._service = new SmoothSailing();
        }

        [TestMethod]
        public void ReverseParantheses()
        {
            // Arrange
            var test1 = "a(bc)de";
            var test2 = "I'm a sentence";
            var test3 = "a(bcdefghijkl(mno)p)q";
            var test4 = "Code(Cha(lle)nge)";

            // Act
            var result1 = this._service.reverseParentheses(test1);
            var result2 = this._service.reverseParentheses(test2);
            var result3 = this._service.reverseParentheses(test3);
            var result4 = this._service.reverseParentheses(test4);

            // Assert
            Assert.AreEqual("acbde", result1);
            Assert.AreEqual("I'm a sentence", result2);
            Assert.AreEqual("apmnolkjihgfedcbq", result3);
            Assert.AreEqual("CodeegnlleahC", result4);
        }
    }
}
